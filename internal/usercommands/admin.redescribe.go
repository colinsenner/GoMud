package usercommands

import (
	"fmt"
	"strings"

	"github.com/volte6/gomud/internal/rooms"
	"github.com/volte6/gomud/internal/templates"
	"github.com/volte6/gomud/internal/users"
	"github.com/volte6/gomud/internal/util"
)

func Redescribe(rest string, user *users.UserRecord, room *rooms.Room) (bool, error) {

	args := util.SplitButRespectQuotes(rest)

	if len(args) < 2 {
		// send some sort of help info?
		infoOutput, _ := templates.Process("admincommands/help/command.rename", nil)
		user.SendText(infoOutput)
		return true, nil
	}

	// Check whether the user has an item in their inventory that matches
	matchItem, found := user.Character.FindInBackpack(args[0])
	rest = strings.Join(args[1:], " ")

	if !found {
		user.SendText(fmt.Sprintf("You don't have a %s to rename.", rest))
	} else {
		// Swap the item location
		user.Character.RemoveItem(matchItem)
		oldNameSimple := matchItem.DisplayName()
		oldName := matchItem.DisplayName()

		matchItem.Redescribe(rest)

		matchItem.Validate()

		user.Character.StoreItem(matchItem)

		user.SendText(
			fmt.Sprintf(`You chant softly and wave your hand over the <ansi fg="item">%s</ansi>. Success! It now has a new description!`, oldNameSimple),
		)
		room.SendText(
			fmt.Sprintf(`<ansi fg="username">%s</ansi> chants softly and waves their hand over <ansi fg="item">%s</ansi>, causing it to glow briefly.`, user.Character.Name, oldName),
			user.UserId,
		)
	}

	return true, nil
}