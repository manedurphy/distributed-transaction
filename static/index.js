let customerData;
const body = document.getElementsByTagName('body');
const first = document.getElementById('firstname');
const last = document.getElementById('lastname');
const signupEmail = document.getElementById('signup-email');
const signupPassword = document.getElementById('signup-password');
const loginEmail = document.getElementById('login-email');
const loginPassword = document.getElementById('login-password');
const signupForm = document.getElementById('signup-form');
const loginForm = document.getElementById('login-form');
const orderForm = document.getElementById('order-form');
const total = document.getElementById('total');

signupForm.addEventListener('submit', function (e) {
    e.preventDefault();
    createCustomer();
});

loginForm.addEventListener('submit', function (e) {
    e.preventDefault();
    login();
});

function createCustomer() {
    const customer = {
        first_name: first.value,
        last_name: last.value,
        email: signupEmail.value,
        password: signupPassword.value,
    };
    fetch('http://localhost:8000/api/customers/v1', {
        method: 'POST',
        body: JSON.stringify(customer),
        headers: {
            'Content-Type': 'application/json',
        },
    })
        .then((resp) => resp.json())
        .then((data) => {
            if (data.customer) {
                customerData = data.customer;
                renderOrderForm();
                return;
            }
            const errMsgDiv = document.createElement('div');
            errMsgDiv.textContent = data.message;
            body[0].prepend(errMsgDiv);
            setTimeout(function () {
                errMsgDiv.remove();
            }, 3000);
        })
        .catch((err) => console.error(err));
}

function login() {
    const req = {
        email: loginEmail.value,
        password: loginPassword.value,
    };
    fetch('http://localhost:8000/api/customers/v1/login', {
        method: 'POST',
        body: JSON.stringify(req),
        headers: {
            'Content-Type': 'application/json',
        },
    })
        .then((resp) => resp.json())
        .then((data) => {
            if (data.customer) {
                customerData = data.customer;
                renderOrderForm();
                return;
            }
            const errMsgDiv = document.createElement('div');
            errMsgDiv.textContent = data.message;
            body[0].prepend(errMsgDiv);
            setTimeout(function () {
                errMsgDiv.remove();
            }, 3000);
        })
        .catch((err) => console.error(err));
}

function renderOrderForm() {
    const walletDiv = document.createElement('div');
    walletDiv.id = 'wallet';
    signupForm.style.display = 'none';
    loginForm.style.display = 'none';
    orderForm.style.display = 'block';
    walletDiv.textContent = `You have ${customerData.wallet} dollars in your wallet`;
    body[0].prepend(walletDiv);
}

function renderLoginForm() {
    signupForm.style.display = 'none';
    loginForm.style.display = 'block';
}

function renderSignupForm() {
    signupForm.style.display = 'block';
    loginForm.style.display = 'none';
}

function placeOrder() {
    const order = {
        customer_id: Number(customerData.id),
        total: Number(total.value),
    };

    fetch('http://localhost:8000/api/orders/v1', {
        method: 'POST',
        body: JSON.stringify(order),
        headers: {
            'Content-Type': 'application/json',
        },
    })
        .then((resp) => resp.json())
        .then((data) => getOrderStatus(data.order))
        .catch((err) => console.error(err));
}

function getOrderStatus({ id, total }) {
    const orderStatusDiv = document.createElement('div');
    const walletDiv = document.getElementById('wallet');
    walletDiv.remove();

    orderForm.style.display = 'none';
    orderStatusDiv.textContent = 'placing order...';
    body[0].appendChild(orderStatusDiv);

    const sse = new EventSource(`http://localhost:8000/sse/${id}`);
    sse.onmessage = function (event) {
        const data = JSON.parse(event.data);

        if (data.status == 'in progress') {
            orderStatusDiv.textContent = 'order placed!';
            customerData.wallet -= total;
            sse.close();

            setTimeout(function () {
                orderStatusDiv.remove();
                renderOrderForm();
            }, 3000);
        }

        if (data.status == 'declined') {
            orderStatusDiv.textContent = 'payment declined';
            sse.close();

            setTimeout(function () {
                orderStatusDiv.remove();
                renderOrderForm();
            }, 3000);
        }
    };
}

function addFundsToWallet() {}
