<script>
    import ItemList from "../components/ItemList.svelte";
    const splittedHash = window.location.hash.split("/");
    const encodedTag = splittedHash[splittedHash.length - 1];
    let tag = decodeURIComponent(encodedTag);
</script>


<svelte:head>
    <title>#{tag}</title>
</svelte:head>
<div id="container">
    <ItemList title={"#" + tag} url={`/api/items-with-tag?tag=${tag}`} notfound="This tag doesn't exists."/>
</div>
