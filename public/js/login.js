document.addEventListener('DOMContentLoaded', () => {
    const form = document.querySelector('#loginForm');
    const btnEntrar = document.querySelector('#submit');
    const mensagem = document.querySelector('#mensagem');

    form.addEventListener('submit', async (e) => {
        e.preventDefault();

        if (btnEntrar.disabled) {
            return;
        }

        try {
            btnEntrar.disabled = true;
            btnEntrar.textContent = 'Entrando...';

            const formData = new FormData(form);
            const data = {
                agencia: formData.get('agencia'),
                conta: formData.get('conta'),
                senha: formData.get('senha')
            };

            const response = await fetch('/api/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            });

            const result = await response.json();

            if (response.ok) {
                mensagem.innerHTML = 'Login realizado com sucesso!';
                console.log('Resposta do servidor:', result);
            } else {
                mensagem.innerHTML = `Erro: ${result.message || 'Login falhou'}`;
            }
        } catch (erro) {
            mensagem.innerHTML = `Erro na requisição: ${erro.message}`;
        } finally {
            btnEntrar.disabled = false;
            btnEntrar.textContent = 'Entrar';
        }
    });
});