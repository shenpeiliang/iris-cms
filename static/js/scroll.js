//页面滚动
 function b(){
			h = 100;
			t = $(document).scrollTop();
			if(t > h){
				$('.r_top').show();
			}else{
				$('.r_top').hide();
			}
		}
$(function(){	
	//回到顶部	
	$(window).scroll(function(e){
		b();		
	})	
	$('.r_top').click(function(){			
		$(document).scrollTop(0);					
	 });
	 //电话号码
	 $('.r_tel').hover(function(){
			$('.tel_info').show();
		},function(){
			$('.tel_info').hide();
	})
});