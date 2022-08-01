<template>
 <div class="page-loader">
    <div class="cube"></div>
    <div class="cube"></div>
    <div class="cube"></div>
    <div class="cube"></div>
  </div>
</template>
<style lang="scss" scoped>
 $colors: #8cc271, #69BEEB, #F5AA39, #E9643B;
 
 .page-loader{
  margin-top: 80px;
  display: flex;
  justify-content: center;
  align-items: center;
 }
 
 .cube{
  width: 32px;
  height: 32px;
  margin-right: 10px;
  border-radius: 50%;
  
  @for $i from 1 through length($colors){
    &:nth-child(#{$i}){
      background-color: nth($colors, $i);
    }
  }
  
  &:first-child{
    animation: left 1s infinite;
  }
  
   &:last-child{
    animation: right 1s infinite .5s;
  }
 }
 
 @keyframes left{
  40%{
    transform: translateX(-60px)
  }
  50%{
     transform: translateX(0)
  }
 }
 
  @keyframes right{
  40%{
    transform: translateX(60px)
  }
  50%{
     transform: translateX(0)
  }
 }
 
 
 
 
</style>