$(document).ready(function(){
    //Carrega Lista de conteudos cadastrados
    $.ajax({
        url: "http://localhost:12345/api/lisContent",
        type:"jsonp",
        crossDomain: true,
        success:function(data) {
            var html = "";
            jQuery.each(JSON.parse(data), function(i, item){
                html += "<tr>";
                html += "<td>"+item.id+"</td>";
                html += "<td>"+item.title+"</td>";
                html += "<td>"+item.description+"</td>";
                html += "<td>"+moment(item.date_ini).format('DD/MM/YYYY')+"</td>";
                html += "<td>"+moment(item.date_end).format('DD/MM/YYYY')+"</td>";
                html += "<td>"+item.group+"</td>";
                html += "<td>"+item.category_content+"</td>";
                html += "<td><a href='/conteudo/"+item.id+"' title='edit'><i class='fa fa-eye text-primary' aria-hidden='true'></i></a>"+
                "<a href='/conteudo/delete/"+item.id+"' title='exclude'><i class='fa fa-trash fa-2 text-danger' aria-hidden='true'></i></a></td>";
                html += "</tr>"; 
            })
            $("#res").html(html)
        }
    });
    //Montar um option para grupos
    $.ajax({
        url: "http://localhost:12345/api/listGroup",
        type:"jsonp",
        crossDomain: true,
        success:function(data) {
            var options = "<option>Selecione um item</option>";
            jQuery.each(JSON.parse(data), function(i, item){
                options += "<option value='"+item.id+"'>"+item.name+"</option>";
            })
            $("#groupsC").html(options)
        }
    });

    $("#groupsC").on("change", function(){
        var id = $("#groupsC").val()
        getCategorias(id);
    });

    var url = window.location.href.split("/")
    var last = url.pop()
    getDadosContent(url[3], last);
})

function getCategorias(grupo) {
    //Montar um option para categorias
    $.ajax({
        url: "http://localhost:12345/api/listCategorysByGroup/"+grupo,
        type:"POST",
        crossDomain: true,
        success:function(data) {
            var options = "<option>Selecione uma categoria</option>";
            jQuery.each(JSON.parse(data), function(i, item){
                options += "<option value='"+item.id+"'>"+item.categoria+"</option>";
            })
            $("#categoriaContent").html(options)
        }
    });
}

function getDadosContent(mod, id) {
    if (mod == "conteudo" && id !="") {
        $.ajax({
            url: "http://localhost:12345/api/listContentByID/"+id,
            type:"jsonp",
            crossDomain: true,
            success:function(result) {
                result = JSON.parse(result);
                //tituloContent, descContent, dateIni, dateEnd, group, categoriaContent, texto
                $("#tituloContent").val(result[0].title);
                $("#descContent").val(result[0].description);
                $("#dateIniC").val(result[0].date_ini);
                $("#dateEndC").val(result[0].date_end);
                $("#groupsC").val(result[0].group);
                $("#categoriaContent").val(result[0].category_content);
                $("#texto").val(result[0].text);

                $.ajax({
                    url: "http://localhost:12345/api/listGroup",
                    type:"jsonp",
                    crossDomain: true,
                    success:function(data) {
                        var selec = "";
                        var options = "<option>Selecione um item</option>";
                        jQuery.each(JSON.parse(data), function(i, item){
                            if (result[0].group == item.id) {
                                selec = 'selected="selected"'
                            } else {
                                selec = ''
                            }
                            options += "<option value='"+item.id+"' "+selec+">"+item.name+"</option>";
                        })
                        $("#groupsC").html(options)
                    }
                });

                $.ajax({
                    url: "http://localhost:12345/api/listCategorysByGroup/"+result[0].group,
                    type:"POST",
                    crossDomain: true,
                    success:function(data) {
                        var selecc = "";
                        var options = "<option>Selecione uma categoria</option>";
                        jQuery.each(JSON.parse(data), function(i, item){
                            if (result[0].category_content == item.id) {
                                selecc = 'selected="selected"'
                            } else {
                                selecc = ''
                            }
                            options += "<option value='"+item.id+"' "+selecc+" >"+item.categoria+"</option>";
                        })
                        $("#categoriaContent").html(options)
                    }
                });
            }
        });
    }
}

function newFunction() {
    return "content";
}
