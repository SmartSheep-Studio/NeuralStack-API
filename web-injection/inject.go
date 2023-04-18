package web

func RenderInjectPages() []map[string]any {
	var rendered []map[string]any
	for _, app := range AppliedApps {
		if len(app.InjectApp) <= 0 {
			continue
		}
		rendered = append(rendered, map[string]any{
			"id":          app.ID,
			"uri":         app.Uri,
			"icon":        app.Icon,
			"name":        app.Name,
			"inject":      app.InjectApp,
			"description": app.Description,
			"x":           app.WindowOptions.X,
			"y":           app.WindowOptions.Y,
			"w":           app.WindowOptions.Width,
			"h":           app.WindowOptions.Height,
			"minW":        app.WindowOptions.MinWidth,
			"minH":        app.WindowOptions.MinHeight,
			"maxW":        app.WindowOptions.MaxWidth,
			"maxH":        app.WindowOptions.MaxHeight,
			"resizeable":  app.WindowOptions.Resizable,
			"closeable":   app.WindowOptions.Closeable,
			"pkg":         app.PackageID,
			"meta":        app.WindowOptions.Meta,
		})
	}
	return rendered
}
