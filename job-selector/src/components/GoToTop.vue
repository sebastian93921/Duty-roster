<template>

<div class="go-top">
    <a class="smoothscroll" title="Back to Top" href="#top"><i class="fa fa-arrow-up" aria-hidden="true"></i></a>
</div>

</template>

<style scoped>
.go-top {
    position: fixed;
    bottom: 5px;
    right: 5px;
    z-index: 600;
    display: none;
}

.go-top a, .go-top a:visited {
    text-decoration: none;
    border: 0 none;
    display: block;
    height: 60px;
    width: 60px;
    line-height: 60px;
    text-align: center;
    background: gray;
    color: rgba(255, 255, 255, 0.5);
    text-align: center;
    text-transform: uppercase;
    -webkit-transition: all 0.3s ease-in-out;
    transition: all 0.3s ease-in-out;
    border-radius: 50%;
    -moz-border-radius: 50px;
    -webkit-border-radius: 50px;
}

.go-top a i, .go-top a:visited i {
    font-size: 18px;
    line-height: inherit;
}

.go-top a:hover, .go-top a:focus {
    color: #FFFFFF;
}
</style>
