var sliding = 0;

// Set is sliding value
function setSliding(a_ISliding){
	sliding = a_ISliding;
}

// Get is sliding value
function getSliding(){
	return sliding;
}

// Carry out accordian styled effect
function accordion(evt) {
	el = Event.element(evt);
	
	//  If element is visible do nothing
	if ($('visible') == el) {
			return;
	}
	if ($('visible')) {
	
			if( getSliding() == 1 ){
					return false;
			}
		
			var eldown = getNextSibling(el);
			var elup = getNextSibling($('visible'));
			
			setSliding( 1 );
			
			new Effect.Parallel(
			[
					new Effect.SlideUp(elup),
					new Effect.SlideDown(eldown)
			], {
					duration: 0.3,
					afterFinish: function() { setSliding( 0 );}
			});

			$('visible').id = '';
	}
	el.id = 'visible';
}


// Setup accordian initial state
function init() {
		
		var bodyPanels = document.getElementsByClassName('panel_body');
		var panels = document.getElementsByClassName('panel');
		var noPanels = panels.length;
		var percentageWidth = 100 / noPanels;
		var position = 0;
		
		//  Loop through body panels and panels applying required styles and adding event listeners
    for (i = 0; i < bodyPanels.length; i++) {
			bodyPanels[i].hide();
			panels[i].style.width = percentageWidth + '%';
			panels[i].style.position = 'absolute';
			panels[i].style.left = position + '%';
			
			Event.observe(panels[i].getElementsByTagName('h3')[0], 'mouseover', accordion, false);
			Event.observe(panels[i].getElementsByTagName('h3')[0], 'mousemove', accordion, false);
			
			position += percentageWidth;
    }
		
		//  Set panel with id of visible to be initial displayed
    var vis = $('visible').parentNode.id+'-body';
    $(vis).show();
}

// Next sibling method to work around firefox issues
function getNextSibling(startBrother){
  endBrother=startBrother.nextSibling;
  while(endBrother.nodeType!=1){
    endBrother = endBrother.nextSibling;
  }
  return endBrother;
}

Event.observe(window, 'load', init, false);