// Aguarda o carregamento completo do DOM (todos os elementos HTML)
document.addEventListener('DOMContentLoaded', () => {
    // Seleciona os elementos do HTML usando suas classes e IDs
    const form = document.querySelector('.form'); // O formulário
    const btnCadastrar = document.querySelector('#submit'); // Botão de envio
    const mensagem = document.querySelector('.mensagem'); // Área de mensagens
    const inputNomeCompleto = document.querySelector('#nomeCompleto'); // Campo de nome completo
    const inputDataNascimento = document.querySelector('#dataNascimento'); // Campo de data de nascimento
    const inputCep = document.querySelector('#cep'); // Campo de cep
    const inputEndereco = document.querySelector('#endereco'); // Campo de endereço
    const inputCpf = document.querySelector('#cpf'); // Campo de cpf
    const inputRg = document.querySelector('#rg'); // Campo de RG
    const inputBairro = document.querySelector('#bairro'); // Campo de Bairro
    const inputCidade = document.querySelector('#cidade'); // Campo de Cidade
    const inputEstado = document.querySelector('#estado'); // Campo de Estado
    const inputRendaMensal = document.querySelector('#rendamensal'); // Campo de renda mensal
    const inputEmail = document.querySelector('#email'); // Campo de e-mail
    const inputSenha = document.querySelector('#senha'); // Campo de senha
    const inputSenhaConfirmada = document.querySelector('senhaConfirmada'); // Campo de senha confirmada

    // Quando a página termina de carregar completamente, coloca o foco no campo de nome completo
    window.addEventListener('load', () => {
        inputNomeCompleto.focus();
    });

    // Quando o campo de nome completo recebe foco (clique), exibe uma mensagem de ajuda
    inputNomeCompleto.addEventListener('focus', () => {
        mensagem.innerHTML = 'Digite o seu nome completo';
    });

    // Quando o campo de e-mail recebe foco (clique), exibe uma mensagem de ajuda
    inputEmail.addEventListener('focus', () => {
        mensagem.innerHTML = 'Digite o seu e-mail no formato email@exemplo.com.br';
    });

    // Quando o campo de data de nascimento recebe foco, exibe uma mensagem de ajuda
    inputDataNascimento.addEventListener('focus', () => {
        mensagem.innerHTML = 'Digite a sua data de nascimento com DD/MM/AAAA';
    });

    // Quando o campo de data de nascimento recebe dados (input), valida o contéudo e exibe uma mensagem de ajuda
    inputDataNascimento.addEventListener('input', () => {
        let value = inputDataNascimento.value.replace(/\D/g, ""); // remove tudo que não for número
        if (value.length > 2) {
            value = value.replace(/^(\d{2})(\d)/, "$1/$2");
        }
        if (value.length > 5) {
            value = value.replace(/^(\d{2})\/(\d{2})(\d)/, "$1/$2/$3");
        }
        inputDataNascimento.value = value;
    });

    // Quando o campo de cpf recebe foco (clique), exibe uma mensagem de ajuda
    inputCpf.addEventListener('focus', () => {
        if (inputCpf.value === "")
            mensagem.innerHTML = 'Digite o seu CPF no formato 999.999.999-99';
        else
            mensagem.innerHTML = 'Cpf inválido';
    });

    // Quando o campo de cpf recebe dados (input), valida o contéudo e exibe uma mensagem de ajuda
    inputCpf.addEventListener('input', () => {
        let value = inputCpf.value.replace(/\D/g, ""); // remove tudo que não for número

        value = value.replace(/(\d{3})(\d)/, "$1.$2");
        value = value.replace(/(\d{3})(\d)/, "$1.$2");
        value = value.replace(/(\d{3})(\d{1,2})$/, "$1-$2");

        inputCpf.value = value;
    });

    inputCpf.addEventListener('blur', () => {
        if (!validarCPF(inputCpf.value)) {
            inputCpf.focus();
            return;
        }
    });

    // Quando o campo de rg recebe foco (clique), exibe uma mensagem de ajuda
    inputRg.addEventListener('focus', () => {
        if (inputRg.value === "")
            mensagem.innerHTML = 'Digite o seu RG no formato 99.999.999-9';
        else
            mensagem.innerHTML = 'RG inválido';
    });

    // Quando o campo de rg recebe dados (input), valida o contéudo e exibe uma mensagem de ajuda
    inputRg.addEventListener('input', () => {
        let value = inputRg.value.replace(/[^0-9Xx]/g, "").toUpperCase();

        if (value.length > 2)
            value = value.replace(/^(\d{2})(\d)/, "$1.$2");

        if (value.length > 6)
            value = value.replace(/^(\d{2})\.(\d{3})(\d)/, "$1.$2.$3");

        if (value.length > 9)
            value = value.replace(/^(\d{2})\.(\d{3})\.(\d{3})([0-9X])$/, "$1.$2.$3-$4");

        inputRg.value = value;
    });

    inputRg.addEventListener('blur', () => {
        const regex = /^\d{2}\.\d{3}\.\d{3}-[0-9Xx]$/;
        if (!regex.test(inputRg.value)) {
            inputRg.focus();
            return;
        }
    });

    // Quando o campo de cep recebe foco, exibe uma mensagem de ajuda
    inputCep.addEventListener('focus', () => {
        if (inputCep.value === "")
            mensagem.innerHTML = 'Digite o cep do seu endereço';
        else
            mensagem.innerHTML = 'Cep inválido, ou inexistente!';
    });

    // Quando o campo de cep recebe dados (input), valida o contéudo e exibe uma mensagem de ajuda
    inputCep.addEventListener('input', () => {
        let value = inputCep.value.replace(/\D/g, ""); // remove tudo que não for número
        inputCep.value = value;
        const regex = /^\d{8}$/;
        if (!regex.test(inputCep.value)) {
            mensagem.innerHTML = 'Cep inválido, deve conter 8 números';
            return;
        } else {
            mensagem.innerHTML = 'Digite o cep da sua rua';
        }
    });

    inputCep.addEventListener('blur', () => {
        fetch(`https://viacep.com.br/ws/${inputCep.value}/json/`)
            .then((response) => {
                console.log(response);
                return response.json();
            })
            .then((json) => {
                console.log(json);
                if (!json.erro) {
                    inputEndereco.value = json.logradouro;
                    inputBairro.value = json.bairro;
                    inputCidade.value = json.localidade;
                    inputEstado.value = json.estado;
                } else {
                    mensagem.innerHTML = 'Cep inválido, ou inexistente!';
                    inputCep.focus();
                    return;
                }
            })
            .catch((erro) => {
                mensagem.innerHTML = 'Cep inválido, ou inexistente!';
                inputCep.focus();
                return;
            });
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

function validarCPF(cpf) {
    cpf = cpf.replace(/\D/g, '');

    if (cpf.length !== 11 || /^(\d)\1+$/.test(cpf)) return false;

    let soma = 0;
    let resto;

    for (let i = 1; i <= 9; i++)
        soma += parseInt(cpf.substring(i - 1, i)) * (11 - i);

    resto = (soma * 10) % 11;
    if (resto === 10 || resto === 11) resto = 0;
    if (resto !== parseInt(cpf.substring(9, 10))) return false;

    soma = 0;
    for (let i = 1; i <= 10; i++)
        soma += parseInt(cpf.substring(i - 1, i)) * (12 - i);

    resto = (soma * 10) % 11;
    if (resto === 10 || resto === 11) resto = 0;

    return resto === parseInt(cpf.substring(10, 11));
}