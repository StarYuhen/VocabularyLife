<view data-event-opts="{{[['tap',[['click',['$event']]]]]}}" class="{{['u-icon','data-v-6fd2a1ec','u-icon--'+labelPos]}}" style="{{$root.s0}}" bindtap="__e"><block qq:if="{{isImg}}"><image class="u-icon__img data-v-6fd2a1ec" style="{{$root.s1}}" src="{{name}}" mode="{{imgMode}}"></image></block><block qq:else><text class="{{['u-icon__icon','data-v-6fd2a1ec',customClass]}}" style="{{$root.s2}}" hover-class="{{hoverClass}}" data-event-opts="{{[['touchstart',[['touchstart',['$event']]]]]}}" bindtouchstart="__e"><block qq:if="{{showDecimalIcon}}"><text class="{{['u-icon__decimal','data-v-6fd2a1ec',decimalIconClass]}}" style="{{$root.s3}}" hover-class="{{hoverClass}}"></text></block></text></block><block qq:if="{{label!==''}}"><text class="u-icon__label data-v-6fd2a1ec" style="{{'color:'+(labelColor)+';'+('font-size:'+($root.g0)+';')+('margin-left:'+(labelPos=='right'?$root.g1:0)+';')+('margin-top:'+(labelPos=='bottom'?$root.g2:0)+';')+('margin-right:'+(labelPos=='left'?$root.g3:0)+';')+('margin-bottom:'+(labelPos=='top'?$root.g4:0)+';')}}">{{label+''}}</text></block></view>