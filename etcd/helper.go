package etcd

import (
	"fmt"
)

func (c *Client) IsKey(path string) (bool, error) {

	resp, err := c.Get(path, false, false)

	if err != nil {
		return false, err
	}

	return !resp.Node.Dir, nil
}

func (c *Client) TryGetValue(path string, value *string) (bool, error) {

	val, err := c.GetValue(path)
	*value = val

	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Client) GetValue(path string) (string, error) {

	resp, err := c.Get(path, false, false)

	if err != nil {
		return "", err
	}

	if resp.Node.Dir {
		return "", fmt.Errorf("provided path % is direcory and no key", path)
	}

	return resp.Node.Value, nil
}

func (c *Client) DirectoryCount(path string) (int, error) {

	resp, err := c.Get(path, false, false)

	if err != nil {
		return 0, err
	}

	if !resp.Node.Dir {
		return 0, fmt.Errorf("provided path % is key and no directory", path)
	}

	nCount := 0
	for _, node := range resp.Node.Nodes {
		if node.Dir { nCount++ }
	}

	return nCount, nil
}
