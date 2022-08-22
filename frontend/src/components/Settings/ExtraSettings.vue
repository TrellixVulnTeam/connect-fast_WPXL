<template>
    <div>
        This will be an extra settings. 
        
        <input type="file" id="recognize" />
        <img src="../../assets/images/troops.png" alt="something"/>
        <button @click="readImage">REad image </button>
    </div>
</template>


<script>
import Tesseract from 'tesseract.js';

export default {
    methods:{
        readImage(){
            Tesseract.recognize(
                "../../assets/images/troops.png",
                'eng', 
                {logger: m => console.log(m)}
                ).then(
                    ({data: {text}}) => {
                        console.log(text)
                    }
                )
        }
    }
}
</script>