$('#register-form').on('submit', createUser);

function createUser(event) {
    event.preventDefault();

    if ($('#password').val() != $('#confirm-password').val()) {
        Swal.fire("Ops...", "As senhas não coincidem!", "error");
        return;
    }

    $.ajax({
        url: "/users",
        method: "POST",
        data: {
           name: $('#name').val(),
           email: $('#email').val(),
           nickname: $('#nickname').val(),
           password: $('#password').val()
        }
    }).done(function() {
        Swal.fire("Sucesso!", "Usuário cadastrado com sucesso!", "success")
            .then(function() {
                $.ajax({
                    url: "/login",
                    method: "POST",
                    data: {
                        email: $('#email').val(),
                        password: $('#password').val()
                    }
                }).done(function() {
                    window.location = "/home";
                }).fail(function() {
                    Swal.fire("Ops...", "Erro ao autenticar o usuário!", "error");
                })
            })
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao cadastrar o usuário!", "error");
    });
}
