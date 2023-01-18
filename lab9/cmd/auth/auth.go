package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"time"

	"code.gitea.io/gitea/modules/auth/pam"
	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func main() {
	log.Fatal(runServer())
}

func runServer() error {
	server := Server{
		secret: []byte(os.Getenv("AUTH_SECRET")),
	}
	router := httprouter.New()
	router.GET("/", server.IndexHandler)
	router.GET("/subscribe", server.SubscribeHandler)
	router.POST("/login", server.LoginHandler)

	s := &http.Server{
		Handler: router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		return err
	}
	log.Printf("listening on http://%v", l.Addr())

	return s.Serve(l)
}

type Server struct {
	secret []byte
}

func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "html/auth.html")
}

func (s *Server) checkToken(token string) (*accessToken, error) {
	var claims jwt.StandardClaims
	parsedToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return s.secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(*jwt.StandardClaims); ok && parsedToken.Valid {
		return &accessToken{
			expires: claims.ExpiresAt,
			user:    claims.Id,
		}, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

type accessToken struct {
	expires int64
	user string
}

func (s *Server) SubscribeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("Created new ws session")
	accessCookie, err := r.Cookie("accessToken")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	claims, err := s.checkToken(accessCookie.Value)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "")

	err = s.subscribe(c, claims.user)
	if websocket.CloseStatus(err) == websocket.StatusNormalClosure ||
		websocket.CloseStatus(err) == websocket.StatusGoingAway {
		return
	}
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *Server) subscribe(c *websocket.Conn, user string) error {
	cmd := exec.Command("sudo", "--user", user, "sh", "-s")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	out := io.MultiReader(stdout, stderr)
	err = cmd.Start()
	if err != nil {
		return err
	}

	errc := make(chan error, 1)
	go func() {
		scanner := bufio.NewScanner(out)
		for scanner.Scan() {
			ctx := context.TODO()
			line := scanner.Text()
			log.Println("sending line of output: ", line)
			err := wsjson.Write(ctx, c, Output{
				Line: line+"\n",
			})
			if err != nil {
				errc <- err
			}
		}
	}()

	go func() {
		var input Input
		for {
			ctx := context.TODO()

			err := wsjson.Read(ctx, c, &input)
			if err != nil {
				errc <- err
			}
			log.Println("received input:", input.Cmd)
			stdin.Write([]byte(input.Cmd+"\n"))
		}
	}()

	return <-errc
}

type Input struct {
	Cmd string `json:"cmd"`
}

type Output struct {
	Line string `json:"line"`
}

type userInfo struct {
	Login string `json:"login"`
	Passwd string `json:"passwd"`
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user userInfo
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	_, err = pam.Auth("passwd", user.Login, user.Passwd)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println(err)
		return
	}

	exp := time.Now().Unix() + 60
	claims := &jwt.StandardClaims{
		ExpiresAt: exp,
		Id:    user.Login,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(s.secret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	log.Println("Logged in as", user.Login)
	http.SetCookie(w, &http.Cookie{
		Name:       "accessToken",
		Value:      ss,
		Expires:    time.Unix(exp, 0),
	})
}
