package rest

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Postcord/objects"
)

func getRealClient(t *testing.T) *Client {
	tokenEnv := os.Getenv("TOKEN")
	if tokenEnv == "" {
		t.Skip("TOKEN env var must be set to run tests against discord API")
	}

	if testing.Short() {
		t.Skip("Test skipped due to short mode")
	}

	return New(&Config{
		Token:       fmt.Sprintf("Bot %s", tokenEnv),
		UserAgent:   fmt.Sprintf("Postcord/1.0 %s (%s)", runtime.GOOS, runtime.GOARCH),
		Ratelimiter: NewMemoryRatelimiter(&MemoryConf{MaxRetries: 3}),
	})
}

func testChannelSnowflake(t *testing.T) objects.Snowflake {
	env := os.Getenv("TEST_CHANNEL_ID")
	if env != "" {
		i, err := strconv.Atoi(env)
		require.NoError(t, err)

		return objects.Snowflake(i)
	}

	return objects.Snowflake(484093378993192971)
}

func TestClient_GetChannel_Integration(t *testing.T) {
	client := getRealClient(t)

	c, err := client.GetChannel(testChannelSnowflake(t))
	if err != nil {
		t.Error("failed to get channel:", err)
	}

	t.Log(c)
}

func TestClient_GetChannelMessages_Integration(t *testing.T) {
	client := getRealClient(t)

	c, err := client.GetChannelMessages(testChannelSnowflake(t), nil)
	if err != nil {
		t.Error("failed to get channel messages:", err)
	}

	t.Log(c)
}
