package main

// func getBASEClaim(ctx context.Context, req events.APIGatewayProxyRequest) (*auth.BASEGuestClaim, error) {
// 	// keys := []string{}
// 	for k, v := range req.Headers {
// 		if strings.ToLower(k) == auth.AuthHeader {
// 			baseClaim, err := auth.VerifyGuestJWT(ctx, v)
// 			if err != nil {
// 				if err.Error() == model.ErrJWTExpire.Error() {
// 					//wlog.LogSystem(ctx, "jwtAuth", wUtil.StrLog(err))
// 					return baseClaim, nil
// 				}
// 				wlog.LogSystem(ctx, "jwtAuth", wUtil.StrLog(err))
// 			}
// 			return baseClaim, err
// 		}
// 	}
// 	return nil, model.ErrAuthHeaderEmpty
// }
