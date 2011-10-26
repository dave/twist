
package twist

func GetTemplateByPath(path string) *Template {
	switch path {
		
		case "html_content_plain1" : 
			return html_content_plain1_Template()
		
		case "html_content_plain2" : 
			return html_content_plain2_Template()
		
		case "html_content_plain3" : 
			return html_content_plain3_Template()
		
		case "html_content_red1" : 
			return html_content_red1_Template()
		
		case "html_content_red2" : 
			return html_content_red2_Template()
		
		case "html_content_red3" : 
			return html_content_red3_Template()
		
		case "html_misc_navigation" : 
			return html_misc_navigation_Template()
		
		case "html_master_plainMaster" : 
			return html_master_plainMaster_Template()
		
		case "html_master_redMaster" : 
			return html_master_redMaster_Template()
		
	}
	return nil
}


func html_content_plain1_Template() *Template{
	return &Template {
		Name     : "plain1",
		Html     : `<script>function template_plain1(id){return "<p>\n	This is a page using the plain_template.html file.\n</p>\n<p>\n	Here's a counter:\n</p>\n<h1 id=\""+id+"_Output\">\n	0\n</h1>\n<p>\n	Click <a href=\"#\" id=\""+id+"_Plus\">plus</a> or <a href=\"#\" id=\""+id+"_Minus\">minus</a>\n</p>\n<input type=\"hidden\" id=\""+id+"_Count\" />"}</script>`,
	}
}

func html_content_plain2_Template() *Template{
	return &Template {
		Name     : "plain2",
		Html     : `<script>function template_plain2(id){return "<p>\n	This is the second plain page.\n</p>"}</script>`,
	}
}

func html_content_plain3_Template() *Template{
	return &Template {
		Name     : "plain3",
		Html     : `<script>function template_plain3(id){return "<p>\n	This is another plain page.\n</p>"}</script>`,
	}
}

func html_content_red1_Template() *Template{
	return &Template {
		Name     : "red1",
		Html     : `<script>function template_red1(id){return "<p>\n	This is a page using the red_template.html file.\n</p>"}</script>`,
	}
}

func html_content_red2_Template() *Template{
	return &Template {
		Name     : "red2",
		Html     : `<script>function template_red2(id){return "<p>\n	This is the second red page.\n</p>"}</script>`,
	}
}

func html_content_red3_Template() *Template{
	return &Template {
		Name     : "red3",
		Html     : `<script>function template_red3(id){return "<p>\n	This is another red page.\n</p>"}</script>`,
	}
}

func html_misc_navigation_Template() *Template{
	return &Template {
		Name     : "navigation",
		Html     : `<script>function template_navigation(id){return "<div style=\"border 1px solid #000000; padding: 5px;\">\n	<a id=\""+id+"_plain1Link\">Plain page 1</a> |\n	<a id=\""+id+"_plain2Link\">Plain page 2</a> |\n	<a id=\""+id+"_plain3Link\">Plain page 3</a> |\n	<a id=\""+id+"_red1Link\">Red page 1</a> |\n	<a id=\""+id+"_red2Link\">Red page 2</a> |\n	<a id=\""+id+"_red3Link\">Red page 3</a>\n</div>"}</script>`,
	}
}

func html_master_plainMaster_Template() *Template{
	return &Template {
		Name     : "plainMaster",
		Html     : `<script>function template_plainMaster(id){return "<div id=\""+id+"_Navigation\" />\n<h1 id=\""+id+"_Header\">Plain page</h1>\n<div id=\""+id+"_Content\" />\n<div id=\""+id+"_Footer\" style=\"color:#ffffff; background-color:#000000; font-weight:bold; padding:5px;\" />"}</script>`,
	}
}

func html_master_redMaster_Template() *Template{
	return &Template {
		Name     : "redMaster",
		Html     : `<script>function template_redMaster(id){return "<div id=\""+id+"_Navigation\" />\n<h1 id=\""+id+"_Header\">Red page</h1>\n<div id=\""+id+"_Topper\" style=\"background-color:#cc0000; color:#ffffff; font-weight:bold; padding:5px;\">\n	<span id=\""+id+"_Location\" /> @ <span id=\""+id+"_Date\" />\n</div>\n<div id=\""+id+"_Content\" />\n<div id=\""+id+"_Footer\" style=\"color:#ffffff; background-color:#cc0000; font-weight:bold; padding:5px;\" />"}</script>`,
	}
}

