package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/Helltale/vk-parser-program/internal/logger"
)

type User struct {
	user_id       string
	fields        string
	user_responce map[string]interface{}
	time          time.Time
}

func NewUser(userid, fields_ string) *User {
	if fields_ == "" {
		return &User{
			time:    time.Now(),
			user_id: userid,
			fields:  "aboutactivities,about,blacklisted,blacklisted_by_me,books,bdate,can_be_invited_group,can_post,can_see_all_posts,can_see_audio,can_send_friend_request,can_write_private_message,career,common_count,connections,contacts,city,crop_photo,domain,education,exports,followers_count,friend_status,has_photo,has_mobile,home_town,photo_100,photo_200,photo_200_orig,photo_400_orig,photo_50,sex,site,schools,screen_name,status,verified,games,interests,is_favorite,is_friend,is_hidden_from_feed,last_seen,maiden_name,military,movies,music,nickname,occupation,online,personal,photo_id,photo_max,photo_max_orig,quotes,relation,relatives,timezone,tv,universities,is_verified",
		}
	}
	return &User{
		time:    time.Now(),
		user_id: userid,
		fields:  fields_,
	}
}

func (u *User) CreateLink(accessToken string, version string) string {
	baseURL := "https://api.vk.com/method/users.get"
	params := url.Values{}
	params.Add("access_token", accessToken)
	params.Add("v", version)
	params.Add("user_ids", u.user_id)
	params.Add("fields", u.fields)

	return fmt.Sprintf("%s?%s", baseURL, params.Encode())
}

func (u *User) Fetch(url string, logger *logger.CombinedLogger) error {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("failed to fetch data", "url", url, "user", u, "err_message", err)
		return fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Error("unexpected status code", "url", url, "user", u, "err_message", err)
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("failed to read response body", "url", url, "user", u, "err_message", err)
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		logger.Error("failed to unmarshal response", "url", url, "user", u, "err_message", err)
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	u.user_responce = responseData
	return nil
}
