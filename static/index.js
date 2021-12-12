let customerData;
const body = document.getElementsByTagName('body')[0];
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
const fundsForm = document.getElementById('funds-form');
const ccList = document.getElementById('cc-list');
const ccForm = document.getElementById('cc-form');
const ccNum = document.getElementById('cc-num');
const expMonth = document.getElementById('exp-month');
const expYear = document.getElementById('exp-year');
const cvv = document.getElementById('cvv');
const walletAdd = document.getElementById('wallet-add');

signupForm.addEventListener('submit', function (e) {
    e.preventDefault();
    createCustomer();
});

loginForm.addEventListener('submit', function (e) {
    e.preventDefault();
    login();
});

ccForm.addEventListener('submit', function (e) {
    e.preventDefault();
    addCreditCard();
});

fundsForm.addEventListener('submit', function (e) {
    e.preventDefault();
    addFundsToWallet();
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
            body.prepend(errMsgDiv);
            setTimeout(function () {
                errMsgDiv.remove();
            }, 3000);
        })
        .catch((err) => console.error(err));
}

function login() {
    const reqBody = {
        email: loginEmail.value,
        password: loginPassword.value,
    };
    fetch('http://localhost:8000/api/customers/v1/login', {
        method: 'POST',
        body: JSON.stringify(reqBody),
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
            body.prepend(errMsgDiv);
            setTimeout(function () {
                errMsgDiv.remove();
            }, 3000);
        })
        .catch((err) => console.error(err));
}

function addCreditCard() {
    const reqBody = {
        customer_id: customerData.id,
        credit_card_number: ccNum.value,
        expiration: `${expYear.value}-${expMonth.value}`,
        cvv: Number(cvv.value),
    };

    fetch('http://localhost:8000/api/payments/v1/credit-card', {
        method: 'POST',
        body: JSON.stringify(reqBody),
        headers: {
            'Content-Type': 'application/json',
        },
    })
        .then((resp) => resp.json())
        .then((data) => {
            if (data.id) {
                renderFundsForm();
                return;
            }
        })
        .catch((err) => console.error(err));
}

function renderOrderForm() {
    const walletDiv = document.createElement('div');
    walletDiv.id = 'wallet';
    signupForm.style.display = 'none';
    loginForm.style.display = 'none';
    fundsForm.style.display = 'none';
    orderForm.style.display = 'block';
    walletDiv.textContent = `You have ${customerData.wallet} dollars in your wallet`;
    body.prepend(walletDiv);
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
    body.appendChild(orderStatusDiv);

    const sse = new EventSource(`http://localhost:8000/sse/${id}`);
    sse.onmessage = function (event) {
        const data = JSON.parse(event.data);

        if (data.status == 'paid') {
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

function addFundsToWallet() {
    const reqBody = {
        customer_id: customerData.id,
        credit_card_id: 1,
        amount: Number(walletAdd.value),
    };

    fetch('http://localhost:8000/api/payments/v1/funds', {
        method: 'POST',
        body: JSON.stringify(reqBody),
    })
        .then((resp) => {
            if (resp.ok) {
                customerData.wallet += reqBody.amount;
            }
            return resp.json();
        })
        .then((data) => {
            if (data.message) {
                const fundsSuccessDiv = document.createElement('div');
                fundsSuccessDiv.textContent = data.message;
                body.prepend(fundsSuccessDiv);
                walletAdd.value = null;

                setTimeout(function () {
                    fundsSuccessDiv.remove();
                }, 3000);
            }
        })
        .catch((err) => console.error(err));
}

function renderCreditCardForm() {
    fundsForm.style.display = 'none';
    ccForm.style.display = 'block';
}

function renderFundsForm() {
    const walletDiv = document.getElementById('wallet');
    if (walletDiv != null) {
        walletDiv.remove();
    }
    fundsForm.style.display = 'block';
    orderForm.style.display = 'none';
    ccForm.style.display = 'none';
    while (ccList.firstChild) {
        ccList.removeChild(ccList.firstChild);
    }

    const reqBody = {
        customer_id: customerData.id,
    };

    fetch('http://localhost:8000/api/payments/v1/credit-cards', {
        method: 'POST',
        body: JSON.stringify(reqBody),
        headers: {
            'Content-Type': 'application/json',
        },
    })
        .then((resp) => resp.json())
        .then((data) => {
            if (data.credit_cards) {
                data.credit_cards.forEach((cc) => {
                    const opt = document.createElement('option');
                    opt.value = cc;
                    opt.textContent = `**** ${cc}`;
                    ccList.appendChild(opt);
                });
                return;
            }

            const opt = document.createElement('option');
            ccList.appendChild(opt);
            opt.textContent = 'No credit cards found';
        })
        .catch((err) => console.error(err));
}
