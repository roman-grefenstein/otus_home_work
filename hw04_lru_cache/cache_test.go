package hw04_lru_cache //nolint:golint,stylecheck

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c, err := NewCache(10)

		require.Nil(t, err)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c, err := NewCache(5)

		require.Nil(t, err)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c, err := NewCache(10)

		require.Nil(t, err)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", "ccc")
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, "ccc", val)

		c.Clear()

		_, ok = c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("not simple", func(t *testing.T) {
		c, err := NewCache(3)

		require.Nil(t, err)

		wasInCache := c.Set("a", "aaa") // {a:ааа}
		require.False(t, wasInCache)

		wasInCache = c.Set("b", "bbb") // {b:bbb a:ааа}
		require.False(t, wasInCache)

		wasInCache = c.Set("c", "ccc") // {c:ccc b:bbb a:ааа}
		require.False(t, wasInCache)

		wasInCache = c.Set("d", "ddd") // {d:ddd c:ccc b:bbb}
		require.False(t, wasInCache)

		val, ok := c.Get("a") // {d:ddd c:ccc b:bbb}
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get("b") // {b:bbb d:ddd c:ccc}
		require.True(t, ok)
		require.Equal(t, "bbb", val)

		wasInCache = c.Set("c", "c") // {c:c b:bbb ddd:400}
		require.True(t, wasInCache)

		wasInCache = c.Set("b", "b") // {b:b c:c ddd:400}
		require.True(t, wasInCache)

		wasInCache = c.Set("a", "a") // {a:а b:b c:c}
		require.False(t, wasInCache)

		val, ok = c.Get("d")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("return error", func(t *testing.T) {
		invalidCapacities := []int{0, -5}

		for _, tc := range invalidCapacities {
			tc := tc
			t.Run(strconv.Itoa(tc), func(t *testing.T) {
				c, err := NewCache(tc)

				require.NotNil(t, err)
				require.Nil(t, c)
			})
		}
	})
}
