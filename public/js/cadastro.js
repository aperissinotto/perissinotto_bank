// Aguarda o carregamento completo do DOM (todos os elementos HTML)
document.addEventListener('DOMContentLoaded', () => {
    // Seleciona os elementos do HTML usando suas classes e IDs
    const form = document.querySelector('.form'); // O formulário
    const btnCadastrar = document.querySelector('#submit'); // Botão de envio
    const mensagem = document.querySelector('.mensagem'); // Área de mensagens
    const inputNomeCompleto = document.querySelector('#nomeCompleto'); // Campo de nome completo
    const inputDataNascimento = document.querySelector('#dataNascimento'); // Campo de data de nascimento
    const inputCep = document.querySelector('#cep'); // Campo de cep

    // Quando a página termina de carregar completamente, coloca o foco no campo de nome completo
    window.addEventListener('load', () => {
        inputNomeCompleto.focus();
    });

    // Quando o campo de nome completo recebe foco (clique), exibe uma mensagem de ajuda
    inputNomeCompleto.addEventListener('focus', () => {
        mensagem.innerHTML = 'Digite o seu nome completo';
    });

    // Quando o campo de data de nascimento recebe foco, exibe uma mensagem de ajuda
    inputDataNascimento.addEventListener('focus', () => {
        mensagem.innerHTML = 'Digite a sua data de nascimento com DD/MM/AAAA';
    });

    // Quando o campo de data de nascimento recebe dados (input), valida o contéudo e exibe uma mensagem de ajuda
    inputDataNascimento.addEventListener('input', () => {
        const regex = /^((0[1-9]|[12][0-9]|3[01]))[/]((0[1-9]|1[012]))[/](\d{4})$/;
        if (!regex.test(inputDataNascimento.value)) {
            mensagem.innerHTML = 'Data de nascimento inválida, deve ser no formato DD(dia)/MM(mês)/AAAA(ano)!';
        } else {
            mensagem.innerHTML = 'Digite a sua data de nascimento com DD/MM/AAAA';
        }
    });

    // Quando o campo de cep recebe foco, exibe uma mensagem de ajuda
    inputCep.addEventListener('focus', () => {
        mensagem.innerHTML = 'Digite o cep da sua rua';
    });

    // Quando o campo de cep recebe dados (input), valida o contéudo e exibe uma mensagem de ajuda
    inputCep.addEventListener('input', () => {
        const regex = /^\d{8}$/;
        if (!regex.test(inputCep.value)) {
            mensagem.innerHTML = 'Cep inválido, pode conter apenas números';
        } else {
            mensagem.innerHTML = 'Digite o cep da sua rua';
        }
    });

    // Quando o formulário é enviado (botão clicado)
    form.addEventListener('submit', async (e) => {
        e.preventDefault(); // Impede o comportamento padrão do formulário

        // Se o botão já está desabilitado, não faz nada (evita múltiplos cliques)
        if (btnCadastrar.disabled) {
            return;
        }

        try {
            // Desabilita o botão e muda o texto para indicar que está processando
            btnCadastrar.disabled = true;
            btnCadastrar.textContent = 'Entrando...';

            // Coleta os dados do formulário
            const formData = new FormData(form);
            const data = {
                nomeCompleto: formData.get('nomeCompleto'),
                dataNascimento: formData.get('dataNascimento'),
                cep: formData.get('cep')
            };

            // Envia os dados para o servidor via requisição POST
            const response = await fetch('/api/buscaCep', {
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
            btnCadastrar.disabled = false;
            btnCadastrar.textContent = 'Entrar';
        }
    });
});