package graphics

type radioSelector struct {
	items []*tappableIcon
}

func (r *radioSelector) Add(icon *tappableIcon) int {
	r.items = append(r.items, icon)
	return len(r.items) - 1
}

func (r *radioSelector) OnItemTap(icon *tappableIcon) {
	for _, item := range r.items {
		if item != icon {
			item.Refresh()
		}
	}
}
