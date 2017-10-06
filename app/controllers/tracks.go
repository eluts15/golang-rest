package controllers

import (
	"api/app/models"

	"github.com/revel/revel"
)

type Tracks struct {
	*revel.Controller
}

var tracks = []models.Track{
	models.Track{1, "Last Hope", "Paramore", "Pop/Rock", "5:37"},
	models.Track{2, "Grow Up", "Paramore", "Pop/Rock", "4:25"},
	models.Track{3, "Looking Up", "Paramore", "Pop/Rock", "4:10"},
}

func (c Tracks) List() revel.Result {
	return c.RenderJSON(tracks)
}

func (c Tracks) Show(trackID int) revel.Result {
	var res models.Track

	for _, track := range tracks {
		if track.ID == trackID {
			res = track
		}
	}

	if res.ID == 0 {
		return c.NotFound("Track not found")
	}

	return c.RenderJSON(res)
}
