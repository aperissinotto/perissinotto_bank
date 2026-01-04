// Aguarda o carregamento completo do DOM (todos os elementos HTML)
document.addEventListener('DOMContentLoaded', () => {
    // Seleciona os elementos do HTML usando suas classes e IDs
    const form = document.querySelector('.form'); // O formulário
    const btnEntrar = document.querySelector('#submit'); // Botão de envio
    const mensagem = document.querySelector('.mensagem'); // Área de mensagens
    const inputCpf = document.querySelector('#cpf'); // Campo de agência
    const inputSenha = document.querySelector('#senha'); // Campo de senha

    // Quando a página termina de carregar completamente, coloca o foco no campo de agência
    window.addEventListener('load', () => {
        inputCpf.focus();
    });

    // Quando o campo de cpf recebe foco (clique), exibe uma mensagem de ajuda
    inputCpf.addEventListener('focus', () => {
        mensagem.innerHTML = 'Digite o número do seu cpf';
    });

    // Quando o campo de cpf recebe dados (input), valida o contéudo e exibe uma mensagem de ajuda
    inputCpf.addEventListener('input', () => {
        let value = inputCpf.value.replace(/\D/g, ""); // remove tudo que não for número
        inputCpf.value = value;
        const regexCpf = /^\d{11}$/;
        if (!regexCpf.test(inputCpf.value)) {
            mensagem.innerHTML = 'CPF inválido!';
        } else {
            mensagem.innerHTML = 'CPF válido';
        }
    });

    // Quando o campo de senha recebe foco, exibe uma mensagem de ajuda
    inputSenha.addEventListener('focus', () => {
        mensagem.innerHTML = 'Digite sua senha de acesso a sua conta';
    });

    // Quando o formulário é enviado (botão clicado)
    form.addEventListener('submit', async (e) => {
        e.preventDefault(); // Impede o comportamento padrão do formulário

        // Valida a senha
        const regexSenha = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/;
        if (!(regexSenha.test(inputSenha.value))) {
            mensagem.innerHTML = 'Digite a sua senha';
            return;
        }

        // Se o botão já está desabilitado, não faz nada (evita múltiplos cliques)
        if (btnEntrar.disabled) {
            return;
        }

        try {
            // Desabilita o botão e muda o texto para indicar que está processando
            btnEntrar.disabled = true;
            btnEntrar.textContent = 'Entrando...';

            // Coleta os dados do formulário
            const formData = new FormData(form);
            const data = {
                cpf: formData.get('cpf'),
                senha: formData.get('senha')
            };

            // Envia os dados para o servidor via requisição POST
            const response = await fetch('/api/login', {
                method: 'POST', // Tipo de requisição
                headers: {
                    'Content-Type': 'application/json' // Indica que está enviando JSON
                },
                body: JSON.stringify(data) // Converte os dados para JSON
            });

            mensagem.innerHTML = 'Validando dados informados no Mainframe...';

            // Aguarda e converte a resposta do servidor em JSON
            const result = await response.json();

            // Se a requisição foi bem-sucedida (código 200, 201, etc)
            if (response.ok) {
                mensagem.innerHTML = 'Login realizado com sucesso!';
            } else {
                // Se houve erro, exibe a mensagem de erro retornada pelo servidor
                mensagem.innerHTML = `Erro: ${result.message || 'Login falhou'}`;
            }
        } catch (erro) {
            // Se houver erro na comunicação com o servidor, exibe a mensagem de erro
            mensagem.innerHTML = `Erro na requisição: ${erro.message}`;
        } finally {
            // Após tudo (sucesso ou erro), reabilita o botão e restaura o texto
            btnEntrar.disabled = false;
            btnEntrar.textContent = 'Entrar';
        }
    });
});