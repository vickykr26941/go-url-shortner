package models

import "time"

type Analytics struct {
	ID          int64     `json:"id" db:"id"`
	URLID       int64     `json:"url_id" db:"url_id"`
	ClickedAt   time.Time `json:"clicked_at" db:"clicked_at"`
	IPAddress   string    `json:"ip_address" db:"ip_address"`
	UserAgent   string    `json:"user_agent" db:"user_agent"`
	Referer     *string   `json:"referer" db:"referer"`
	Country     *string   `json:"country" db:"country"`
	City        *string   `json:"city" db:"city"`
	Browser     *string   `json:"browser" db:"browser"`
	OS          *string   `json:"os" db:"os"`
	DeviceType  *string   `json:"device_type" db:"device_type"`
	UTMSource   *string   `json:"utm_source" db:"utm_source"`
	UTMMedium   *string   `json:"utm_medium" db:"utm_medium"`
	UTMCampaign *string   `json:"utm_campaign" db:"utm_campaign"`
}

func TableName() string {
	return "analytics"
}

type AnalyticsRequest struct {
	IPAddress   string  `json:"ip_address"`
	UserAgent   string  `json:"user_agent"`
	Referer     *string `json:"referer"`
	UTMSource   *string `json:"utm_source"`
	UTMMedium   *string `json:"utm_medium"`
	UTMCampaign *string `json:"utm_campaign"`
}

type AnalyticsSummary struct {
	URLID             int64                 `json:"url_id"`
	TotalClicks       int64                 `json:"total_clicks"`
	UniqueVisitors    int64                 `json:"unique_visitors"`
	ClicksByDate      []DailyClickStats     `json:"clicks_by_date"`
	ClicksByCountry   []CountryClickStats   `json:"clicks_by_country"`
	ClicksByBrowser   []BrowserClickStats   `json:"clicks_by_browser"`
	ClicksByOS        []OSClickStats        `json:"clicks_by_os"`
	ClicksByDevice    []DeviceClickStats    `json:"clicks_by_device"`
	ClicksByReferer   []RefererClickStats   `json:"clicks_by_referer"`
	ClicksByUTMSource []UTMSourceClickStats `json:"clicks_by_utm_source"`
}

type DailyClickStats struct {
	Date   string `json:"date"`
	Clicks int64  `json:"clicks"`
}

type CountryClickStats struct {
	Country string `json:"country"`
	Clicks  int64  `json:"clicks"`
}

type BrowserClickStats struct {
	Browser string `json:"browser"`
	Clicks  int64  `json:"clicks"`
}

type OSClickStats struct {
	OS     string `json:"os"`
	Clicks int64  `json:"clicks"`
}

type DeviceClickStats struct {
	DeviceType string `json:"device_type"`
	Clicks     int64  `json:"clicks"`
}

type RefererClickStats struct {
	Referer string `json:"referer"`
	Clicks  int64  `json:"clicks"`
}

type UTMSourceClickStats struct {
	UTMSource string `json:"utm_source"`
	Clicks    int64  `json:"clicks"`
}

type AnalyticsDateRange struct {
	StartDate time.Time `json:"start_date" query:"start_date"`
	EndDate   time.Time `json:"end_date" query:"end_date"`
}
