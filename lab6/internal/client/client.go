package client

import (
	"fmt"
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
	conn, err := ftp.Dial(os.Getenv("HOST"), ftp.DialWithTimeout(5*time.Second))
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
	entries, err := c.conn.List(path)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.Type == ftp.EntryTypeFolder {
			fmt.Printf("%s/\n", entry.Name)
		} else {
			fmt.Println(entry.Name)
		}
	}
	return nil
}
