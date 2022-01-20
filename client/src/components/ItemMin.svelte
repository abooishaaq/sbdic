<script>
    export let item = {};

    let {title, updated_at, count} = item;

    let updatedAt = new Date(updated_at).toDateString();
</script>

<a href={`#/item/${encodeURIComponent(title)}`}>
    <div class="item">
        <p class="item-title">{title}</p>
        <p>count: {count}</p>
        <p>updated at: {updatedAt}</p>
    </div>
</a>


<style>
    a {
        display: block;
        width: 100%;
    }
    .item {
        color: #000;
        display: grid;
        grid-template-columns: 3fr 1fr 2fr;
    }
    .item-title {
        font-size: 1.2em;
    }
</style>
