function getValues(items) {
	$.each(items, function(i,n){try{items[i].V = $("#" + items[i].I).val()}catch(ex){}})
}