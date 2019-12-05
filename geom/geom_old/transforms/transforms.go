package transforms

import (
	"github.com/tadeuszjt/geom"
)

func Cam2DToCamera(cameraRect, displayRect geom.Rect, cameraOrientation geom.Ori2) geom.Mat33 {
	// scale
	sx := cameraRect.Width() / displayRect.Width()
	sy := cameraRect.Height() / displayRect.Height()

	// translate
	tx := cameraRect.Min.X - sx*displayRect.Min.X
	ty := cameraRect.Min.Y - sy*displayRect.Min.Y

	// align between rectangles
	align := geom.Mat33{
		{sx, 0, tx},
		{0, sy, ty},
		{0, 0, 1},
	}

	return cameraOrientation.Mat33Transform().Times(align)
}

func Cam2DToDisplay(cameraRect, displayRect geom.Rect, cameraOrientation geom.Ori2) geom.Mat33 {
	// scale
	sx := cameraRect.Width() / displayRect.Width()
	sy := cameraRect.Height() / displayRect.Height()

	// translate
	tx := cameraRect.Min.X - sx*displayRect.Min.X
	ty := cameraRect.Min.Y - sy*displayRect.Min.Y

	// align between rectagles
	align := geom.Mat33{
		{1 / sx, 0, -tx},
		{0, 1 / sy, -ty},
		{0, 0, 1},
	}

	return align.Times(cameraOrientation.Mat33InvTransform())
}
