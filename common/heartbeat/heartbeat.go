package heartbeat

import ()

type Args struct {

}

type Response struct {
	Res string
}

type HeartbeatServer struct {

}

func (hb *HeartbeatServer) SendHealthcheck(args *Args, res *Response) error {
	res.Res = "response string"
	return nil
}