var Checkout = function () {};

Checkout.prototype = {
    constructor: Checkout,
    basketId: '',
    
    updatePrices: function () {
        var product = $('#product');
        var price = $('#price');

        product.on('change', function () {
            $('#quantity').val(1);
            switch (product.val()) {
                case 'VOUCHER':
                    price.val('€ 5.00');
                    return;
                case 'TSHIRT':
                    price.val('€ 20.00');
                    return;
                case 'MUG':
                    price.val('€ 7.50');
                    return;
                default:
                    price.val('');
                    return;
            }
        });
    },

    newCheckout: function(checkout) {
        $('#newCheckout').on('click', function () {
            checkout.requestNewBasket(checkout);
            $('#newCheckout').hide();
            $('.basket').show();
        });
    },

    addProduct: function(checkout) {
        $('#addProduct').on('click', function () {
            checkout.requestAddProduct(checkout)
        })
    },

    requestAddProduct: function(checkout) {
        console.log(checkout.basketId);

        return $.ajax({
            url: 'http://localhost:3001/v1/checkout/basket/' + checkout.basketId + '/products',
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(
                {
                    'product-code': $('#product').val(),
                    'quantity': $('#quantity').val()
                }
            ),
            context: this
        }).done(function (data) {
            console.log(data);
            checkout.requestAmount(checkout)
        }).fail(function (response) {
            console.log(response);
            alert('Something goes wrong, sorry, try again later.');
        });
    },

    requestNewBasket: function(checkout) {
        return $.ajax({
            url: 'http://localhost:3001/v1/checkout/basket',
            method: 'POST',
            contentType: 'application/json',
            context: this
        }).done(function (data) {
            checkout.basketId = data.id;
        }).fail(function (response) {
            console.log(response);
            alert('Something goes wrong, sorry, try again later.');
        });
    },

    requestAmount: function(checkout) {
        return $.ajax({
            url: 'http://localhost:3001/v1/checkout/basket/' + checkout.basketId + '/amount',
            method: 'GET',
            contentType: 'application/json',
            context: this
        }).done(function (data) {
            var amount = parseFloat(data.amount).toFixed(2);
            $('.amount').text('Total: € ' + amount);
            $('.amount').css('font-size', '1.5rem');
            $('.amount').css('margin-bottom', '.5rem');
            $('.amount').css('font-weight', '500');
            $('.amount').css('line-height', '1.2');
        }).fail(function (response) {
            console.log(response);
            alert('Something goes wrong, sorry, try again later.');
        });
    },

    registerEvents: function (checkout) {
        checkout.updatePrices();
        checkout.newCheckout(checkout);
        checkout.addProduct(checkout);
    }
};

$(document).ready(function() {
    var checkout = new Checkout();
    checkout.registerEvents(checkout);
});
