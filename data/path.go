package data

const AUTH_PREFIX = "/auth"
const LOGIN_PATH = AUTH_PREFIX + "/login"
const REFRESH_TOKEN_PATH = AUTH_PREFIX + "/refresh-token"
const LOGOUT_PATH = AUTH_PREFIX + "/logout"
const REGISTER_PATH = AUTH_PREFIX + "/register"
const FORGOT_PASSWORD_PATH = AUTH_PREFIX + "/forgot-password"
const RESET_PASSWORD_PATH = AUTH_PREFIX + "/reset-password"

const USER_PREFIX = "/user"
const MY_PROFILE_PATH = USER_PREFIX + "/me"
const USER_UPDATE_PROFILE_PATH = USER_PREFIX + "/update"

const MoviePrefix = "/movie"
const MovieListPath = MoviePrefix + "/list"
const MovieCreatePath = MoviePrefix + "/create"
const MovieUpdatePath = MoviePrefix + "/update"
const MovieDeletePath = MoviePrefix + "/delete"
const MovieUploadPath = MoviePrefix + "/upload"
const LikedMoviePath = MoviePrefix + "/liked"
const WatchedMoviePath = MoviePrefix + "/watched"
const PlayMoviesPath = MoviePrefix + "/play"
