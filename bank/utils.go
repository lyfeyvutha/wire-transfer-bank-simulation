package main

func (c *bankServer) hasEnoughCash(usrID userID, target int32) bool {
	return c.userToCash[usrID] >= target
}

func (c *bankServer) hasEnoughStocks(usrID userID, stocks int32) bool {
	return c.userToStocks[usrID] >= stocks
}
