package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/chanxuehong/wechat/menu"
	"io/ioutil"
)

// 创建自定义菜单.
//  NOTE: 创建自定义菜单后，由于微信客户端缓存，需要24小时微信客户端才会展现出来。
//  建议测试时可以尝试取消关注公众账号后再次关注，则可以看到创建后的效果。
func (c *Client) MenuCreate(mn *menu.Menu) error {
	if mn == nil {
		return errors.New("menu == nil")
	}
	token, err := c.Token()
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(mn)
	if err != nil {
		return err
	}

	_url := clientMenuCreateURL(token)
	resp, err := c.httpClient.Post(_url, postJSONContentType, bytes.NewReader(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result Error
	if err = json.Unmarshal(body, &result); err != nil {
		return err
	}
	if result.ErrCode != 0 {
		return &result
	}
	return nil
}

// 删除自定义菜单
func (c *Client) MenuDelete() error {
	token, err := c.Token()
	if err != nil {
		return err
	}

	_url := clientMenuDeleteURL(token)
	resp, err := c.httpClient.Get(_url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result Error
	if err = json.Unmarshal(body, &result); err != nil {
		return err
	}
	if result.ErrCode != 0 {
		return &result
	}
	return nil
}

// 获取自定义菜单
func (c *Client) MenuGet() (*menu.Menu, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}

	_url := clientMenuGetURL(token)
	resp, err := c.httpClient.Get(_url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Menu menu.Menu `json:"menu"`
		Error
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	return &result.Menu, nil
}