package utils

func GetStructByUri() map[string]interface{} {
	return map[string]interface{}{
		"/sign-up":                new(SignupSchema),
		"/login":                  new(LoginSchema),
		"/user/password/reset":    new(ResetPasswordSchema),
		"/tweet":                  new(TweetSchema),
		"/comment":                new(CommentSchema),
		"/user":                   new(EditSchema),
		"/user/password":          new(ChangePassword),
		"/notification/subscribe": new(Subscription),
	}
}
