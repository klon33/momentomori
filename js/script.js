(function ($) {

    $('.button2').click(function (e) {

        $.post("/post.php", {}, function (result) {
            $(".result").html(result);
        });
        return false;
    });


})(window.jQuery);