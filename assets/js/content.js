/**
 * Inicio do onload
 */
$(document).ready(function(){
    //Carrega Lista de conteudos cadastrados
    $.ajax({
        url: "/api/lisContent",
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
                html += "<td>";
                html += "<button class='btn btn-primary' onclick='updateContent("+item.id+")' title='exclude'><i class='fa fa-eye fa-2 text-secundary' aria-hidden='true'></i></button>&nbsp;";
                html += "<button class='btn btn-danger' onclick='excludeContent("+item.id+")' title='exclude'><i class='fa fa-trash fa-2 text-secundary' aria-hidden='true'></i></button>";
                html += "</td>";
                html += "</tr>"; 
            })
            $("#res").html(html)
        }
    });
    //Montar um option para grupos
    $.ajax({
        url: "/api/listGroup",
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
});//FIM DO ONLOAD
/**
 * Retorna todas as categorias referente a um grupo
 * @param {*codigo do grupo} grupo 
 */
function getCategorias(grupo) {
    //Montar um option para categorias
    $.ajax({
        url: "/api/listCategorysByGroup/"+grupo,
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
/**
 * Retorna as informações do conteudo
 * @param {*modulo no qual estamos acessando} mod 
 * @param {*codigo referente ao modulo} id 
 */
function getDadosContent(mod, id) {
    if (mod == "conteudo" && id !="") {
        $.ajax({
            url: "/api/listContentByID/"+id,
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
                $("#codigo").val(result[0].id);

                $.ajax({
                    url: "/api/listGroup",
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
                    url: "/api/listCategorysByGroup/"+result[0].group,
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

                listFileContent(result[0].id);
            }
        });
    }
}
/**
 * Abre template para carregar alterar um conteudo
 * @param {*} id 
 */
function updateContent(id) {
    setTimeout(function(){
        window.location= "/conteudo/"+id;
    },100)
}
/**
 * Monta tabela com todos os arquivos do conteudo
 * @param {*} id 
 */
function listFileContent(id) {
    $.ajax({
        url: "/api/listFileContent/"+id,
        type:"POST",
        crossDomain: true,
        success:function(data) {
            var html = "";
            jQuery.each(JSON.parse(data), function(i, item){
                html += "<tr>";
                html += "<td><a href='/static/"+item.path+"' download><img src='/static/"+item.path+"' width='50' height='50'></a></td>";
                html += "<td>"+item.nome+"</td>";
                html += "<td><input name='coment' id='"+item.id+"' value='"+item.comentario+"' onchange='saveComent("+item.id+",this.value)'></td>";
                html += "<td>";
                html += "<button class='btn btn-danger' onclick='excludeFile("+item.id+", "+id+")' title='exclude'><i class='fa fa-trash fa-2 text-secundary' aria-hidden='true'></i></button>";
                html += "</td>";
                html += "</tr>"; 
            })
            $("#files").html(html)
        }
    });
}
/**
 * Exclui um arquivo relacioanado a um conteudo e volta para pagina corrata
 * @param {*id arquivo} id 
 * @param {* id do conteudo mãe} content 
 */
function excludeFile(id, content) {
    $.ajax({
        url: "/api/deleteComent/"+id,
        type:"PST",
        success:function(data) {
            jQuery.each(JSON.parse(data), function(i, item){
                if (item.status === 302)  {
                    swal('Alterado!',item.menssage,'success');
                    setTimeout(function(){
                        window.location= "/conteudo/"+content;
                    },1000)
                } else {
                    swal('Erro!',item.menssage,'error');
                }
            });
        }
    });
}
/**
 * Salva um comentarios para foto
 * @param {*codigo do comentario} id 
 * @param {*valor a ser inserido} value 
 */
function saveComent(id, value) {
    $.ajax({
        url: "/api/saveComent/",
        type:"GET",
        data: {cod:id, valor:value},
        success:function(data) {
            jQuery.each(JSON.parse(data), function(i, item){
                if (item.status === 302)  {
                    swal('Alterado!',item.menssage,'success');
                } else {
                    swal('Erro!',item.menssage,'error');
                }
            });
        }
    });

}
/**
 * Excluindo um conteudo
 * @param {*int} id 
 */
function excludeContent(id) {
    swal({
        title: 'DELETE?',
        text: "VOCÊ REAMENTE DESEJA EXCLUIR ESSA INFORMAÇÃO?",
        type: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        cancelButtonText: 'NÃO! ',
        confirmButtonText: 'SIM! ',
        confirmButtonClass: 'btn btn-success',
        cancelButtonClass: 'btn btn-danger',
        buttonsStyling: false,
        reverseButtons: true
      }).then((result) => {
        if (result.value) {
            deleteC("deleteContent/"+id);
            // result.dismiss can be 'cancel', 'overlay',
            // 'close', and 'timer'
        } else if (result.dismiss === 'cancel') {
            swal(
                'Cancelled',
                'Your imaginary file is safe :)',
                'error'
            )
        }
    })
}
/**
 * Execução em segundo plano
 * @param {*} params 
 */
function deleteC(params) {
    $.ajax({
        url: "/api/"+params,
        type:"post",
        crossDomain: true,
        success:function(data) {
            jQuery.each(JSON.parse(data), function(i, item){
                if (item.status == 302) {
                    swal(
                        'Deleted!',
                        'Conteudo Excluido.',
                        'success'
                        )
                        setTimeout(function() {
                            window.location='/conteudo'
                        }, 1000)
                } else {
                    swal(
                        'Deleted!',
                        'Conteudo não Excluido.',
                        'error'
                    )
                }
            })
        }
    });
}