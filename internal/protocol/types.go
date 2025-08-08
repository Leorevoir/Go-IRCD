package protocol

/*
numeric reply codes as defined in RFC 1459
*/
const (
	RPL_WELCOME     = 001
	RPL_YOURHOST    = 002
	RPL_CREATED     = 003
	RPL_MYINFO      = 004
	RPL_TOPIC       = 332
	RPL_NAMREPLY    = 353
	RPL_ENDOFNAMES  = 366
	RPL_MOTDSTART   = 375
	RPL_MOTD        = 372
	RPL_ENDOFMOTD   = 376
	RPL_WHOREPLY    = 352
	RPL_ENDOFWHO    = 315
	RPL_WHOISUSER   = 311
	RPL_WHOISSERVER = 312
	RPL_ENDOFWHOIS  = 318
	RPL_LISTSTART   = 321
	RPL_LIST        = 322
	RPL_LISTEND     = 323
	RPL_NOTOPIC     = 331

	ERR_NOSUCHNICK      = 401
	ERR_NOSUCHCHANNEL   = 403
	ERR_UNKNOWNCOMMAND  = 421
	ERR_NONICKNAMEGIVEN = 431
	ERR_NOTREGISTERED   = 451
	ERR_NEEDMOREPARAMS  = 461
	ERR_NICKNAMEINUSE   = 433
)

/*
* represents an IRC message
 */
type Message struct {
	Prefix   string
	Command  string
	Params   []string
	Trailing string
}
