package statics

import (
	"strings"
	"path/filepath"
	//"fmt"
)


// shameless stolen from OneForAll project
var blockedLibs = []string{
	"npm.js",
    "bower.js",
    "component.js",
    "spm.js",
    "jam.js",
    "jspm.js",
    "ender.js",
    "volo.js",
    "duo.js",
    "yarn.js",
    "requirejs.js",
    "browserify.js",
    "seajs.js",
    "headjs.js",
    "curl.js",
    "lazyload.js",
    "systemjs.js",
    "lodjs.js",
    "esl.js",
    "modulejs.js",
    "webpack.js",
    "rollup.js",
    "brunch.js",
    "parcel.js",
    "microbundle.js",
    "typescript.js",
    "hegel.js",
    "typl.js",
    "mocha.js",
    "jasmine.js",
    "qunit.js",
    "jest.js",
    "prova.js",
    "dalekjs.js",
    "protractor.js",
    "tape.js",
    "testcafe.js",
    "ava.js",
    "cypress.js",
    "chai.js",
    "enzyme.js",
    "proxyquire.js",
    "istanbul.js",
    "blanket.js",
    "jscover.js",
    "phantomjs.js",
    "slimerjs.js",
    "casperjs.js",
    "zombie.js",
    "totoro.js",
    "karma.js",
    "nightwatch.js",
    "intern.js",
    "yolpo.js",
    "puppeteer.js",
    "webdriverio.js",
    "prettier.js",
    "jshint.js",
    "jscs.js",
    "jsfmt.js",
    "jsinspect.js",
    "eslint.js",
    "jslint.js",
    "aurelia.js",
    "backbone.js",
    "meteor.js",
    "ractive.js",
    "vue.js",
    "svelte.js",
    "knockout.js",
    "spine.js",
    "canjs.js",
    "react.js",
    "hyperapp.js",
    "preact.js",
    "nativescript.js",
    "riot.js",
    "thorax.js",
    "chaplin.js",
    "marionette.js",
    "ripple.js",
    "rivets.js",
    "derby.js",
    "jsblocks.js",
    "liquidlava.js",
    "feathers.js",
    "keo.js",
    "atvjs.js",
    "makefun.js",
    "keystonejs.js",
    "ghost.js",
    "apostrophe.js",
    "taracotjs.js",
    "nodizecms.js",
    "cody.js",
    "pencilblue.js",
    "strapi.js",
    "factor.js",
    "nunjucks.js",
    "dot.js",
    "dustjs.js",
    "eco.js",
    "pug.js",
    "ejs.js",
    "xtemplate.js",
    "marko.js",
    "swig.js",
    "ehtml.js",
    "d3.js",
    "peity.js",
    "raphael.js",
    "echarts.js",
    "vis.js",
    "arbor.js",
    "cubism.js",
    "vega.js",
    "envisionjs.js",
    "rickshaw.js",
    "flot.js",
    "nvd3.js",
    "trianglify.js",
    "d4.js",
    "epoch.js",
    "c3.js",
    "babylonjs.js",
    "recharts.js",
    "graphicsjs.js",
    "mxgraph.js",
    "amchart.js",
    "anychart.js",
    "plotly.js",
    "highchart.js",
    "handsontable.js",
    "ace.js",
    "codemirror.js",
    "esprima.js",
    "quill.js",
    "pen.js",
    "editor.js",
    "epiceditor.js",
    "jsoneditor.js",
    "squire.js",
    "tinymce.js",
    "trix.js",
    "trumbowyg.js",
    "wysihtml5.js",
    "popline.js",
    "summernote.js",
    "devdocs.js",
    "dexy.js",
    "docco.js",
    "styledocco.js",
    "ronn.js",
    "dox.js",
    "jsdox.js",
    "esdoc.js",
    "yuidoc.js",
    "coddoc.js",
    "sphinx.js",
    "jsduck.js",
    "codecrumbs.js",
    "jbinary.js",
    "diff2html.js",
    "jspdf.js",
    "underscore.js",
    "lodash.js",
    "sugar.js",
    "ramda.js",
    "mout.js",
    "mesh.js",
    "preludejs.js",
    "rxjs.js",
    "bacon.js",
    "kefir.js",
    "highland.js",
    "mobx.js",
    "mori.js",
    "buckets.js",
    "hashmap.js",
    "moment.js",
    "date.js",
    "fecha.js",
    "dayjs.js",
    "voca.js",
    "selecting.js",
    "he.js",
    "multiline.js",
    "jsurl.js",
    "plexis.js",
    "odometer.js",
    "localforage.js",
    "jstorage.js",
    "cookies.js",
    "crumbsjs.js",
    "randomcolor.js",
    "color.js",
    "colors.js",
    "pleasejs.js",
    "tinycolor.js",
    "i18next.js",
    "polyglot.js",
    "babelfish.js",
    "ttag.js",
    "async.js",
    "q.js",
    "step.js",
    "contra.js",
    "bluebird.js",
    "when.js",
    "objecteventtarget.js",
    "sporadic.js",
    "director.js",
    "pathjs.js",
    "crossroads.js",
    "navaid.js",
    "dompurify.js",
    "log.js",
    "conzole.js",
    "loglevel.js",
    "minilog.js",
    "storyboard.js",
    "regex101.js",
    "regexr.js",
    "regexpbuilder.js",
    "annyang.js",
    "axios.js",
    "bottleneck.js",
    "amygdala.js",
    "wretch.js",
    "farfetch.js",
    "tailor.js",
    "convnetjs.js",
    "dn2a.js",
    "synapses.js",
    "bowser.js",
    "matcha.js",
    "prismjs.js",
    "nprogress.js",
    "pace.js",
    "topbar.js",
    "nanobar.js",
    "pageloadingeffects.js",
    "spinkit.js",
    "ladda.js",
    "ajaxload.js",
    "preloaders.js",
    "cssload.js",
    "validatr.js",
    "formvalidation.js",
    "fieldval.js",
    "funval.js",
    "mousetrap.js",
    "keymaster.js",
    "keypress.js",
    "keyboardjs.js",
    "jwerty.js",
    "shepherd.js",
    "tourist.js",
    "pageguide.js",
    "hopscotch.js",
    "joyride.js",
    "focusable.js",
    "izitoast.js",
    "messenger.js",
    "noty.js",
    "pnotify.js",
    "toastr.js",
    "notie.js",
    "swiper.js",
    "slick.js",
    "slidesjs.js",
    "flexslider.js",
    "unslider.js",
    "sly.js",
    "vegas.js",
    "sequence.js",
    "strut.js",
    "photoswipe.js",
    "jcslider.js",
    "slidr.js",
    "flickity.js",
    "jqrangeslider.js",
    "nouislider.js",
    "fancyinput.js",
    "awesomplete.js",
    "pikaday.js",
    "fullcalendar.js",
    "rome.js",
    "datedropper.js",
    "select2.js",
    "chosen.js",
    "dropzone.js",
    "fileapi.js",
    "plupload.js",
    "form.js",
    "countable.js",
    "card.js",
    "stretchy.js",
    "analytics.js",
    "tipsy.js",
    "opentip.js",
    "qtip2.js",
    "tooltipster.js",
    "simptip.js",
    "toolbar.js",
    "vex.js",
    "sweetalert.js",
    "colorbox.js",
    "fancybox.js",
    "swipebox.js",
    "jbox.js",
    "scrollmonitor.js",
    "headroom.js",
    "iscroll.js",
    "skrollr.js",
    "parallax.js",
    "plax.js",
    "jparallax.js",
    "fullpage.js",
    "scrollmenu.js",
    "simpleparallax.js",
    "slideout.js",
    "jtable.js",
    "datatables.js",
    "tabulator.js",
    "floatthead.js",
    "masonry.js",
    "packery.js",
    "isotope.js",
    "flexboxgrid.js",
    "w2ui.js",
    "fluidity.js",
    "ink.js",
    "dataformsjs.js",
    "webplate.js",
    "cerberus.js",
    "touchemulator.js",
    "dragula.js",
    "leaflet.js",
    "cesium.js",
    "gmaps.js",
    "polymaps.js",
    "jqvmap.js",
    "openlayers3.js",
    "html5media.js",
    "polyplayer.js",
    "flowplayer.js",
    "mediaelement.js",
    "soundjs.js",
    "clappr.js",
    "exifr.js",
    "bigtext.js",
    "circletype.js",
    "slabtext.js",
    "velocity.js",
    "transitionend.js",
    "textillate.js",
    "animatable.js",
    "tsparticles.js",
    "pica.js",
    "cropper.js",
    "es6features.js",
    "gridsome.js",
    "docusaurus.js",
    "echo.js",
    "picturefill.js",
    "json3.js",
    "mixitup.js",
    "grid.js",
    "ky.js",
    "fcal.js",
    "iooxa.js",
    "idyll.js",
    "primer.js",
    "glue.js",
    "postcss.js",
    "mui.js",
    "img2css.js",
    "weui.js",
    "csscss.js",
    "simditor.js",
    "htmlhint.js",
    "csslint.js",
    "grunt.js",
    "yeoman.js",
    "gulp.js",
    "zrender.js",
    "highcharts.js",
    "tweenjs.js",
    "swipe.js",
    "superslides.js",
    "slider.js",
    "polymer.js",
    "ionic.js",
    "timelinejs.js",
    "togetherjs.js",
    "foundation.js",
    "todomvc.js",
    "vuejs.js",
    "webuploader.js",
    "fastclick.js",
    "wangeditor.js",
    "tooling.js",
    "judge.js",
    "amdoc.js",
    "amazeui.js",
    "zepto.js",
    "nodeclub.js",
    "nodeppt.js",
    "hexo.js",
    "koa.js",
    "connect.js",
    "nvm.js",
    "flux.js",
    "browserquest.js",
    "html5shiv.js",
    "ulkit.js",
    "arttemplate.js",
    "jade.js",
    "modernizr.js",
    "css3please.js",
    "babel.js",
    "f2etest.js",
    "brackets.js",
    "ueditor.js",
    "electron.js",
    "base.js",
    "basscss.js",
    "bootflat.js",
    "bootswatch.js",
    "bulma.js",
    "cardinal.js",
    "caramel.js",
    "corpus.js",
    "kube.js",
    "materialize.js",
    "milligram.js",
    "papercss.js",
    "papier.js",
    "pavilion.js",
    "picnicss.js",
    "pure.js",
    "skeleton.js",
    "tachyons.js",
    "tacit.js",
    "uikit.js",
    "wing.js",
    "angular.js",
    "choo.js",
    "deku.js",
    "displayjs.js",
    "inferno.js",
    "mercury.js",
    "mithril.js",
    "moon.js",
    "skatejs.js",
    "solid.js",
    "bliss.js",
    "cash.js",
    "jquery.js",
    "nanojs.js",
    "selector.js",
    "umbrella.js",
    "zeptojs.js",
    "chartist.js",
    "charts.js",
    "chartjs.js",
    "dc.js",
    "dimple.js",
    "d3xter.js",
    "f2.js",
    "frappe.js",
    "ggraph.js",
    "jsplumb.js",
    "metricsgraphics.js",
    "morrisjs.js",
    "muze.js",
    "sparkline.js",
    "sparky.js",
    "taucharts.js",
    "uvcharts.js",
    "vivagraph.js",
    "z3d.js",
    "kartograph.js",
    "mapsicon.js",
    "osmbuildings.js",
    "planetary.js",
    "smallworld.js",
    "tangram.js",
    "topojson.js",
    "turf.js",
    "dynatables.js",
    "listjs.js",
    "sortable.js",
    "tablesaw.js",
    "flipside.js",
    "nudge.js",
    "glide.js",
    "lory.js",
    "siema.js",
    "blotter.js",
    "fitty.js",
    "flowtype.js",
    "lettering.js",
    "shave.js",
    "typeplate.js",
    "fitvid.js",
    "medialementjs.js",
    "plyr.js",
    "talkie.js",
    "videojs.js",
    "abcjs.js",
    "audio5js.js",
    "bap.js",
    "blip.js",
    "howler.js",
    "soundcite.js",
    "tonal.js",
    "vexflow.js",
    "easeljs.js",
    "konva.js",
    "panzoom.js",
    "p5js.js",
    "vizflow.js",
    "zdog.js",
    "scenejs.js",
    "whitestormjs.js",
    "camanjs.js",
    "grafijs.js",
    "smartcrop.js",
    "basicscroll.js",
    "moveto.js",
    "scrollmagic.js",
    "scrollreveal.js",
    "verge.js",
    "alloyfinger.js",
    "anime.js",
    "choreographer.js",
    "gsap.js",
    "impulse.js",
    "mojs.js",
    "popmotion.js",
    "rebound.js",
    "repaintless.js",
    "shifty.js",
    "snabbt.js",
    "snapsvg.js",
    "vivus.js",
    "dotjs.js",
    "handlebars.js",
    "hogan.js",
    "mustache.js",
    "vdo.js",
    "aja.js",
    "fetch.js",
    "qwest.js",
    "reqwest.js",
    "superagent.js",
    "bean.js",
    "eventemitter2.js",
    "mitt.js",
    "elegant.js",
    "feather.js",
    "flaticon.js",
    "fontawesome.js",
    "fontello.js",
    "icomoon.js",
    "ikonate.js",
    "ionicons.js",
    "octicons.js",
    "weloveiconfonts.js",
    "chromajs.js",
    "coolors.js",
    "colorbrewer2.js",
    "colorhexa.js",
    "colourco.js",
    "colormind.js",
    "kewler.js",
    "khroma.js",
    "polychrome.js",
    "uigradients.js",
    "forerunnerdb.js",
    "lokijs.js",
    "lovefield.js",
    "pouchdb.js",
    "rxdb.js",
    "taffydb.js",
    "zangodb.js",
    "parsley.js",
    "dateformat.js",
    "flatpickr.js",
    "instadate.js",
    "luxon.js",
    "tinytime.js",
    "l10ns.js",
    "globalize.js",
    "datakit.js",
    "datalib.js",
    "gauss.js",
    "jstat.js",
    "statkit.js",
    "stdlib.js",
    "theoremjs.js",
    "stealjs.js",
    "aload.js",
    "lazysizes.js",
    "loadxt.js",
    "unveil.js",
    "brain.js",
    "mind.js",
    "neurojs.js",
    "rrssb.js",
    "sharingbuttons.js",
    "socialcount.js",
    "moutjs.js",
    "ramdajs.js",
    "formstone.js",
    "tether.js",
    "upup.js",
    "iconfont.js",
    "iconfont2.js",
    "jwplayer.js",
    "polyfills.js",
    "md5.js",
    "mifihybrid.js"}

//Exist desc
func Exist(str string)bool{	
	for _, item := range blockedLibs {
		alt:=alterations(item)
		for _,v:=range alt{
			if v == str {
				return true
			}
		}		
	}

	return false
}

func alterations(file string)[]string{

	var res []string
	var alterations = []string{"min"}
	
	fulName:=strings.TrimSuffix(file, filepath.Ext(file))


	for _, alt := range alterations {
		if(len(alt) < 1){
			continue
		}
		
		strSplit := strings.Split(fulName, ".")

		for i := 0; i <= len(strSplit); i++ {
			val:=insert(strSplit, i, alt)
			newName:= strings.Join(val, ".") + filepath.Ext(file)
			//fmt.Printf("NEWNMAE: %v\n", newName)
			res = append(res, newName)
		}
	}

	return res
}	


func insert(a []string, index int, value string) []string {
    if len(a) == index { 
        return append(a, value)
    }
    a = append(a[:index+1], a[index:]...)
    a[index] = value
    return a
}