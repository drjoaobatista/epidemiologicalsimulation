line := charts.NewLine()
line.SetGlobalOptions(charts.TitleOpts{Title: "Line-平滑曲线"})
line.AddXAxis(nameItems).AddYAxis("商家A", randInt(), charts.LineOpts{Smooth: true})
