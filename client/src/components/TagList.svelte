<script>
    export let tags = [];

    let textContent;
    const render = (event) => {
        const splitted =  textContent.split(" ");
        if(event.key === " ") {
            tags = [];
            textContent = "";
            console.log(splitted)
            splitted.forEach(x => {
                tags.push(x)
            });
        }
    }
</script>

<ul>
    {#each tags as tag}
        <li>{tag}</li>
    {/each}
</ul>

<style>
    li {
        list-style: none;
        background-color: #dedede;
        display: inline;
    }
</style>