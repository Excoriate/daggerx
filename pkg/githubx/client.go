package githubx

type Config interface {
	GetToken() string
	GetOwner() string
	GetRepo() string
}

type Client struct {
	owner string
	repo  string
	token string
}

func NewClientConfig(owner, repo, token string) Config {
	return &Client{
		owner: owner,
		repo:  repo,
		token: token,
	}
}

func (c *Client) GetOwner() string {
	return c.owner
}

func (c *Client) GetRepo() string {
	return c.repo
}

func (c *Client) GetToken() string {
	return c.token
}

type GHClient struct {
	cfg Config
}

func New(cfg Config) *GHClient {
	return &GHClient{cfg: cfg}
}
