package dto

type CreateUserFollowingByIdParams struct {
	Id string `uri:"id" binding:"required"`
}

type GetUserFollowersByIdParams struct {
	Id string `uri:"id" binding:"required"`
}

type GetUserFollowingByIdParams struct {
	Id string `uri:"id" binding:"required"`
}

type DeleteUserFollowingByIdParams struct {
	Id string `uri:"id" binding:"required"`
}
