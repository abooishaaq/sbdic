<script>
    let text = "loading";
    setInterval(() => {
        if(text.length == 10 ) text = "loading"
        else text += ".";
    }, 300);
</script>


<p>{text}</p>
