package graphqlModel

import "github.com/shurcooL/graphql"

type LoginMutation struct {
	Login struct {
		ID             graphql.String `graphql:"_id"`
		Email          graphql.String
		EmailVerified  graphql.Boolean
		Username       graphql.String
		Nickname       graphql.String
		Company        graphql.String
		Photo          graphql.String
		Browser        graphql.String
		Token          graphql.String
		TokenExpiredAt graphql.String
		LoginsCount    graphql.Int
		LastLogin      graphql.String
		SignedUp       graphql.String
		Blocked        graphql.Boolean
		IsDeleted      graphql.Boolean
	} `graphql:"login(registerInClient: $registerInClient, username: $username, password: $password, email: $email, phone: $phone, phoneCode: $phoneCode)"`
}
