// モーダルを開く
function openModal(bookId) {
    const modal = document.getElementById("myModal");
    const recommendList = document.getElementById("recommendList");

    // 推奨リンクをAPIから取得する
    fetch(`/api/recommend/${bookId}`)
        .then(response => response.json())
        .then(data => {
            recommendList.innerHTML = ""; // リストをクリア
            if (data.recommendations.length > 0) {
                data.recommendations.forEach(recommend => {
                    const listItem = document.createElement("li");
                    const link = document.createElement("a");
                    link.href = recommend.Url;
                    link.textContent = recommend.Detail;
                    link.target = "_blank";
                    listItem.appendChild(link);
                    recommendList.appendChild(listItem);
                });
            } else {
                recommendList.innerHTML = "<li>推奨リンクはありません。</li>";
            }
        });

    modal.style.display = "block";
}

// モーダルを閉じる
function closeModal() {
    const modal = document.getElementById("myModal");
    modal.style.display = "none";
}

// モーダルの外をクリックしたときに閉じる
window.onclick = function(event) {
    const modal = document.getElementById("myModal");
    if (event.target == modal) {
        modal.style.display = "none";
    }
};