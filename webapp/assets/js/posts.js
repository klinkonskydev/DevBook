$('#new-post').on('submit', createPost);

$(document).on('click', '.upvote-post', upvotePost);
$(document).on('click', '.downvote-post', downvotePost);

$('#update-post').on('click', updatePost);
$('.delete-post').on('click', deletePost);

function createPost(event) {
    event.preventDefault();

    $.ajax({
        url: "/posts",
        method: "POST",
        data: {
            title: $('#title').val(),
            content: $('#content').val(),
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao criar a publicação!", "error");
    })
}

function upvotePost(event) {
    event.preventDefault();

    const clickedElement = $(event.target);
    const postId = clickedElement.closest('div').data('post-id');

    clickedElement.prop('disabled', true);
    $.ajax({
        url: `/posts/${postId}/upvote`,
        method: "POST"
    }).done(function() {
        const upvoteCounter = clickedElement.next('span');
        const upvoteCount = parseInt(upvoteCounter.text());

        upvoteCounter.text(upvoteCount + 1);

        clickedElement.addClass('downvote-post');
        clickedElement.addClass('text-danger');
        clickedElement.removeClass('upvote-post');

    }).fail(function() {
        Swal.fire("Ops...", "Erro ao curtir a publicação!", "error");
    }).always(function() {
        clickedElement.prop('disabled', false);
    });
}

function downvotePost(event) {
    event.preventDefault();

    const clickedElement = $(event.target);
    const postId = clickedElement.closest('div').data('post-id');

    clickedElement.prop('disabled', true);
    $.ajax({
        url: `/posts/${postId}/downvote`,
        method: "POST"
    }).done(function() {
        const upvoteCounter = clickedElement.next('span');
        const upvoteCount = parseInt(upvoteCounter.text());

        upvoteCounter.text(upvoteCount - 1);

        clickedElement.removeClass('downvote-post');
        clickedElement.removeClass('text-danger');
        clickedElement.addClass('upvote-post');

    }).fail(function() {
        Swal.fire("Ops...", "Erro ao descurtir a publicação!", "error");
    }).always(function() {
        clickedElement.prop('disabled', false);
    });
}

function updatePost() {
    $(this).prop('disabled', true);

    const postId = $(this).data('post-id');

    $.ajax({
        url: `/posts/${postId}`,
        method: "PUT",
        data: {
            title: $('#title').val(),
            content: $('#content').val()
        }
    }).done(function() {
        Swal.fire('Sucesso!', 'Publicação criada com sucesso!', 'success')
            .then(function() {
                window.location = "/home";
            })
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao editar a publicação!", "error");
    }).always(function() {
        $('#update-post').prop('disabled', false);
    })
}

function deletePost(event) {
    event.preventDefault();

    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir essa publicação? Essa ação é irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmation) {
        if (!confirmation.value) return;

        const clickedElement = $(event.target);
        const post = clickedElement.closest('div')
        const postId = post.data('post-id');

        clickedElement.prop('disabled', true);

        $.ajax({
            url: `/posts/${postId}`,
            method: "DELETE"
        }).done(function() {
            post.fadeOut("slow", function() {
                $(this).remove();
            });
        }).fail(function() {
            Swal.fire("Ops...", "Erro ao excluir a publicação!", "error");
        });
    })

}
