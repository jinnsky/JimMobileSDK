package jimsdk

func (c *Client) SendLogout() {
  c.removeCookieJar()
}