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
    const inputRendaMensal = document.querySelector('#rendaMensal'); // Campo de renda mensal
    const inputEmail = document.querySelector('#email'); // Campo de e-mail
    const inputSenha = document.querySelector('#senha'); // Campo de senha
    const inputSenhaConfirmada = document.querySelector('#senhaConfirmada'); // Campo de senha confirmada
    let camposPreValidados = 0;

    const FLAGS = {
        nomeCompletoValido: 1 << 0,
        dataNascimentoValida: 1 << 1,
        cepValido: 1 << 2,
        enderecoValido: 1 << 3,
        cpfValido: 1 << 4,
        rgValido: 1 << 5,
        rendaMensalValida: 1 << 6,
        emailValido: 1 << 7,
        senhaValida: 1 << 8,
        senhaConfirmadaValida: 1 << 9
    };

    const CAMPOS_OBRIGATORIOS =
        FLAGS.cepValido |
        FLAGS.cpfValido |
        FLAGS.dataNascimentoValida |
        FLAGS.emailValido |
        FLAGS.enderecoValido |
        FLAGS.nomeCompletoValido |
        FLAGS.rendaMensalValida |
        FLAGS.rgValido |
        FLAGS.senhaConfirmadaValida |
        FLAGS.senhaValida;

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

    // Quando o campo de data de nascimento recebe dados (input), ajuda com preenchimento automatico
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
        mensagem.innerHTML = 'Digite o seu CPF no formato 99999999999';
    });

    // Quando o campo de cpf recebe dados (input), ajuda com preenchimento automatico
    inputCpf.addEventListener('input', () => {
        let value = inputCpf.value.replace(/\D/g, ""); // remove tudo que não for número
        inputCpf.value = value;
    });

    // Quando o campo de rg recebe foco (clique), exibe uma mensagem de ajuda
    inputRg.addEventListener('focus', () => {
        mensagem.innerHTML = 'Digite o seu RG no formato 999999999';
    });

    // Quando o campo de rg recebe dados (input), ajuda com preenchimento automatico
    inputRg.addEventListener('input', () => {
        let value = inputRg.value.replace(/[^0-9Xx]/g, "").toUpperCase();
        inputRg.value = value;
    });

    // Quando o campo de cep recebe foco, exibe uma mensagem de ajuda
    inputCep.addEventListener('focus', () => {
        mensagem.innerHTML = 'Digite o cep do seu endereço';
    });

    // Quando o campo de cep recebe dados (input), valida o contéudo e exibe uma mensagem de ajuda
    inputCep.addEventListener('input', () => {
        let value = inputCep.value.replace(/\D/g, ""); // remove tudo que não for número
        inputCep.value = value;
        const regexCep = /^\d{8}$/;
        if (regexCep.test(inputCep.value)) {
            fetch(`https://viacep.com.br/ws/${inputCep.value}/json/`)
                .then((response) => {
                    return response.json();
                })
                .then((json) => {
                    console.log(json);
                    if (!json.erro) {
                        inputEndereco.value = json.logradouro;
                        inputEndereco.value = inputEndereco.value.toUpperCase();
                        inputBairro.value = json.bairro;
                        inputBairro.value = inputBairro.value.toUpperCase();
                        inputCidade.value = json.localidade;
                        inputCidade.value = inputCidade.value.toUpperCase();
                        inputEstado.value = json.estado;
                        inputEstado.value = inputEstado.value.toUpperCase();
                        camposPreValidados |= FLAGS.cepValido;
                    } else {
                        mensagem.innerHTML = 'Cep inválido, ou inexistente!';
                        camposPreValidados &= ~FLAGS.cepValido;
                    }
                })
                .catch((erro) => {
                    mensagem.innerHTML = 'Cep inválido, ou inexistente!';
                    camposPreValidados &= ~FLAGS.cepValido;
                });
        } else {
            console.log('Cep inválido');
            mensagem.innerHTML = 'Cep inválido, deve conter 8 números';
            camposPreValidados &= ~FLAGS.cepValido;
        }
    });

    inputEndereco.addEventListener('focus', () => {
        mensagem.innerHTML = 'Digite o número da residência';
    });

    inputSenha.addEventListener('focus', () => {
        mensagem.innerHTML = 'Digite a senha desejada com no mínimo de 8 caracteres, sendo pelo menos 1 especial e 1 numérico';
    });

    inputSenhaConfirmada.addEventListener('focus', () => {
        mensagem.innerHTML = 'Repita a mesma senha para confirmar';
    });

    inputRendaMensal.addEventListener('focus', () => {
        mensagem.innerHTML = 'Digite sua renda mensal';
    });

    // Quando o formulário é enviado (botão clicado)
    form.addEventListener('submit', async (e) => {
        e.preventDefault(); // Impede o comportamento padrão do formulário

        // Valida o nome completo
        const regexNomeCompleto = /^.{30,80}$/;
        if (regexNomeCompleto.test(inputNomeCompleto.value)) {
            inputNomeCompleto.value = inputNomeCompleto.value.toUpperCase();
            camposPreValidados |= FLAGS.nomeCompletoValido;
        } else {
            mensagem.innerHTML = 'Digite o seu nome completo entre 30 e 80 caracteres';
            camposPreValidados &= ~FLAGS.nomeCompletoValido;
        }

        // Valida o e-mail
        const regexEmail = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
        if (regexEmail.test(inputEmail.value)) {
            inputEmail.value = inputEmail.value.toUpperCase();
            camposPreValidados |= FLAGS.emailValido;
        } else {
            mensagem.innerHTML = 'Digite o seu e-mail no formato email@exemplo.com.br';
            camposPreValidados &= ~FLAGS.emailValido;
        }

        // Valida a data de nascimento
        const regexDataNascimento = /^(0[1-9]|[12][0-9]|3[01])\/(0[1-9]|1[0-2])\/\d{4}$/;
        if (regexDataNascimento.test(inputDataNascimento.value)) {
            camposPreValidados |= FLAGS.dataNascimentoValida;
        } else {
            mensagem.innerHTML = 'Digite a sua data de nascimento com DD/MM/AAAA';
            camposPreValidados &= ~FLAGS.dataNascimentoValida;
        }

        // Valida o CPF
        if (validarCPF(inputCpf.value)) {
            camposPreValidados |= FLAGS.cpfValido;
        } else {
            mensagem.innerHTML = 'Digite o seu CPF no formato 99999999999';
            camposPreValidados &= ~FLAGS.cpfValido;
        }

        // Valida o RG
        const regexRg = /^\d{8}[0-9A-Za-z]$/;
        if (regexRg.test(inputRg.value)) {
            camposPreValidados |= FLAGS.rgValido;
        } else {
            mensagem.innerHTML = 'Digite o seu RG no formato 999999999';
            camposPreValidados &= ~FLAGS.rgValido;
        }

        // Valida o endereço com número
        const regexEndereco = /(^|[^\w])\d+([^\w]|$)/;
        if (regexEndereco.test(inputEndereco.value)) {
            camposPreValidados |= FLAGS.enderecoValido;
        } else {
            mensagem.innerHTML = 'Digite o número da residência';
            camposPreValidados &= ~FLAGS.enderecoValido;
        }

        // Valida a senha
        const regexSenha = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/;
        if (regexSenha.test(inputSenha.value)) {
            camposPreValidados |= FLAGS.senhaValida;
        } else {
            mensagem.innerHTML = 'Digite a senha desejada com no mínimo de 8 caracteres, sendo pelo menos 1 especial e 1 numérico';
            camposPreValidados &= ~FLAGS.senhaValida;
        }

        // Valida a senha confirmada
        if (inputSenhaConfirmada.value === inputSenha.value) {
            camposPreValidados |= FLAGS.senhaConfirmadaValida;
        } else {
            mensagem.innerHTML = 'Senhas inválidas, são diferentes';
            camposPreValidados &= ~FLAGS.senhaConfirmadaValida;
        }

        // Valida a renda mensal
        const regexRendaMensal = /^\d+(\.\d{1,2})?$/;
        if (regexRendaMensal.test(inputRendaMensal.value)) {
            camposPreValidados |= FLAGS.rendaMensalValida;
        } else {
            mensagem.innerHTML = 'Renda inválida';
            camposPreValidados &= ~FLAGS.rendaMensalValida;
        }


        // verifica se todos os campos foram validados antes de submeter
        if ((camposPreValidados & CAMPOS_OBRIGATORIOS) !== CAMPOS_OBRIGATORIOS) {
            console.log('Campos Inválidos!');
            mensagem.innerHTML = 'Existem campos inválidos no formulário';
            return;
        } else {
            console.log('Campos Válidos!');
        }

        // Se o botão já está desabilitado, não faz nada (evita múltiplos cliques)
        if (btnCadastrar.disabled) {
            return;
        }

        try {
            // Desabilita o botão e muda o texto para indicar que está processando
            btnCadastrar.disabled = true;
            btnCadastrar.textContent = 'Cadastrando...';

            // Coleta os dados do formulário
            const formData = new FormData(form);
            const data = {
                nomeCompleto: formData.get('nomeCompleto'),
                email: formData.get('email'),
                dataNascimento: formData.get('dataNascimento'),
                cpf: formData.get('cpf'),
                rg: formData.get('rg'),
                cep: formData.get('cep'),
                endereco: formData.get('endereco'),
                bairro: formData.get('bairro'),
                cidade: formData.get('cidade'),
                estado: formData.get('estado'),
                rendaMensal: Number(formData.get('rendaMensal'))
            };

            // Envia os dados para o servidor via requisição POST
            const response = await fetch('/api/clientes', {
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
                mensagem.innerHTML = 'Cadastro realizado com sucesso!';
            } else {
                // Se houve erro, exibe a mensagem de erro retornada pelo servidor
                mensagem.innerHTML = `Erro: ${result.message || 'Cadastro falhou'}`;
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

// Algoritmo para validar o CPF
function validarCPF(cpf) {
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