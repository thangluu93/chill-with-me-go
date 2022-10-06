package data

const AuthPrefix = "/auth"
const LoginPath = AuthPrefix + "/login"
const RefreshTokenPath = AuthPrefix + "/refresh-token"
const LogoutPath = AuthPrefix + "/logout"
const RegisterPath = AuthPrefix + "/register"
const ForgotPasswordPath = AuthPrefix + "/forgot-password"
const ResetPasswordPath = AuthPrefix + "/reset-password"

const UserPrefix = "/user"
const MyProfilePath = UserPrefix + "/me"
const UserUpdateProfilePath = UserPrefix + "/update"

const MoviePrefix = "/movie"
const MovieListPath = MoviePrefix + "/list"
const MovieCreatePath = MoviePrefix + "/create"
const MovieUpdatePath = MoviePrefix + "/update"
const MovieDeletePath = MoviePrefix + "/delete"
const MovieUploadPath = MoviePrefix + "/upload"
const LikedMoviePath = MoviePrefix + "/liked"
const WatchedMoviePath = MoviePrefix + "/watched"
const PlayMoviesPath = MoviePrefix + "/play"
