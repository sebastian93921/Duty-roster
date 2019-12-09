<template>
<!-- ==================================================
footer
================================================== -->
<footer>

    <div class="footer-bottom">
        <div class="copyright">
            <span>StarsMurasaki</span>
            <span>strsx.com</span>
            <span>© Copyright Sebastian Ko 2018</span>
        </div>


    </div> <!-- end footer-bottom -->

    <div class="go-to-top">
        <GoToTop/>
    </div>

</footer> <!-- end footer -->
</template>

<script>
import GoToTop from './components/GoToTop'

export default {
  name: 'Footer',
  components: {
    GoToTop
  }
}
</script>

<style scoped>
footer {
    padding: 0.2rem 0 0.6rem;
    font-family: "montserrat-light", sans-serif;
    font-size: 15px;
    line-height: 2;
    text-align: center;
    color: rgba(0, 0, 0, 0.3);
    background: #fff
}

footer h4 {
    color: #FFFFFF;
    font-size: 1.6rem;
    line-height: 1.313;
    margin-top: 0;
    margin-bottom: 3.6rem;
}

footer a {
    color: #FFFFFF;
}

footer a:hover, footer a:focus, footer a:active {
    color: #39e0ff;
}

/* ---------------------------------------------------------------
 * ## footer main
 * --------------------------------------------------------------- */
.footer-main {
    text-align: left;
}

.footer-main .left {
    padding-right: 40px;
}

.footer-main .right {
    padding-left: 40px;
}

/* ---------------------------------------------------------------
 * ## footer bottom
 * --------------------------------------------------------------- */
.footer-bottom {
    margin-top: 2rem;
    text-align: center;
    font-size: 14px;
}

.footer-main .copyright span {
    display: inline-block;
}

.footer-bottom .copyright span::after {
    content: "|";
    display: inline-block;
    padding: 0 1rem 0 1.2rem;
    color: rgba(0, 0, 0, 0.2);
}

.footer-bottom .copyright span:last-child::after {
    display: none;
}

/* -------------------------------------------------------------------
 * responsive:
 * footer
 * ------------------------------------------------------------------- */
@media only screen and (max-width: 900px) {
    .footer-main .left {
        padding-right: 20px;
    }
    .footer-main .right {
        padding-left: 20px;
    }
}

@media only screen and (max-width: 800px) {
    footer .row {
        max-width: 600px;
    }
    .footer-main .left {
        padding-right: 15px;
        margin-bottom: 4.2rem;
    }
    .footer-main .right {
        padding-left: 15px;
    }
}

@media only screen and (max-width: 600px) {
    .footer-main .left {
        padding-right: 10px;
    }
    .footer-main .right {
        padding-left: 10px;
    }
    .footer-logo {
        margin: -.6rem auto 3.6rem;
    }
    .footer-bottom {
        padding-bottom: .6rem;
    }
    .footer-bottom .copyright span {
        display: block;
    }
    .footer-bottom .copyright span::after {
        display: none;
    }
}

@media only screen and (max-width: 500px) {
    .footer-subscribe #mc-form input[type="email"] {
        padding: 1.2rem 20px 1.2rem;
        background: rgba(0, 0, 0, 0.1);
        text-align: center;
    }
    .footer-subscribe #mc-form input[type=submit] {
        position: static;
        width: 100%;
        margin-bottom: 1.8rem;
    }
    .footer-subscribe #mc-form label {
        text-align: center;
    }
}

@media only screen and (max-width: 400px) {
    .footer-main .left {
        padding-right: 0;
    }
    .footer-main .right {
        padding-left: 0;
    }
}


</style>
