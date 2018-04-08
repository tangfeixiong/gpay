package server

func (s *Server) CreateCrd(ctx context.Context, req *pb.CrdReqResp) (*pb.CrdReqResp, error) {
	fmt.Println("Request to create CRD:", req)
	resp := &pb.CrdReqResp{
		Recipe: new(pb.CrdRecipient),
	}
	if req == nil || req.Recipe == nil {
		resp.StateCode = 100
		resp.StateMessage = "CRD recipe is required"
		return resp, errors.New(resp.StateMessage)
	}
	if req.Recipe.Plural == "" || req.Recipe.Group == "" {
		resp.StateCode = 101
		resp.StateMessage = "Empty CRD field is not allowed"
		return resp, errors.New(resp.StateMessage)
	}

	resp.Recipe.Group = req.Recipe.Group
	resp.Recipe.Version = req.Recipe.Version
	resp.Recipe.Scope = req.Recipe.Scope
	resp.Recipe.Plural = req.Recipe.Plural
	resp.Recipe.Singular = req.Recipe.Singular
	resp.Recipe.Kind = req.Recipe.Kind
	err := s.ops["rabbitmq-operator"].CreateCRD(req.Recipe)
	if err != nil {
		glog.Infof("Create CRD failed: %s", err.Error())
		resp.StateCode = 10
		resp.StateMessage = err.Error()
		return resp, err
	}
	return resp, nil
}

func (s *Server) ReapCrd(ctx context.Context, req *pb.CrdReqResp) (*pb.CrdReqResp, error) {
	fmt.Println("Requesting:", req)
	return new(pb.CrdReqResp), nil
}
