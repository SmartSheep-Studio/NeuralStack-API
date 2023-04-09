package web

func RenderWindowRegistry() []map[string]any {
	var rendered []map[string]any
	for _, app := range AppliedApps {
		rendered = append(rendered, map[string]any{
			"id":           app.ID,
			"uri":          app.Uri,
			"icon":         app.Icon,
			"name":         app.Name,
			"descriptions": app.Descriptions,
			"x":            app.WindowOptions.X,
			"y":            app.WindowOptions.Y,
			"w":            app.WindowOptions.Width,
			"h":            app.WindowOptions.Height,
			"minW":         app.WindowOptions.MinWidth,
			"minH":         app.WindowOptions.MinHeight,
			"maxW":         app.WindowOptions.MaxWidth,
			"maxH":         app.WindowOptions.MaxHeight,
			"resizeable":   app.WindowOptions.Resizable,
			"closeable":    app.WindowOptions.Closeable,
			"pkg":          app.PackageID,
			"meta":         app.WindowOptions.Meta,
		})
	}
	return rendered
}
