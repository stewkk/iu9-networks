package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/jlaffaye/ftp"
)

type Client struct {
	conn *ftp.ServerConn
}

func NewClient() (Client, error) {
	conn, err := ftp.Dial(os.Getenv("HOST"), ftp.DialWithTimeout(5*time.Second), ftp.DialWithDisabledEPSV(false))
	if err != nil {
		return Client{}, err
	}

	err = conn.Login(os.Getenv("LOGIN"), os.Getenv("PASSWD"))
	return Client{
		conn: conn,
	}, err
}

func (c *Client) Store(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	err = c.conn.Stor(path, f)
	return err
}

func (c *Client) Get(remotePath string) error {
	r, err := c.conn.Retr(remotePath)
	if err != nil {
		return err
	}
	defer r.Close()

	data , err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Base(remotePath), data, 0644)
	return err
}

func (c *Client) Mkdir(path string) error {
	return c.conn.MakeDir(path)
}

func (c *Client) Remove(path string) error {
	err := c.conn.Delete(path)
	if err != nil {
		err = c.conn.RemoveDirRecur(path)
	}
	return err
}

func (c *Client) Ls(path string) error {
	entries, err := c.Nlist(path)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		fmt.Println(entry)
	}
	return nil
}

func (c *Client) Nlist(path string) ([]string, error) {
	return c.conn.NameList(path)
}

func (c *Client) Read(path string) (io.ReadCloser, error) {
	return c.conn.Retr(path)
}

func (c *Client) Load(data io.Reader, path string) error {
	return c.conn.Stor(path, data)
}

func (c *Client) Cwd(path string) error {
	return c.conn.ChangeDir(path)
}
