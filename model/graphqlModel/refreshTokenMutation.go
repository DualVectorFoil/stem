package graphqlModel

import "github.com/kelvinji2009/graphql"

type RefreshTokenMutation struct {
	RefreshToken struct {
		Token graphql.String
		Iat   graphql.Int
		Exp   graphql.Int
	} `graphqlModel:"refreshToken(client: $client, user: $user)"`
}
