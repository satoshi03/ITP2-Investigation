function makeRequestWithUserGesture() {
  var promise = document.requestStorageAccess();
  promise.then(
    function () {
      // Storage access was granted.
      // Check whether the user is logged in.
      // If not, do a popup to log the user in.
      console.log("storage access was granted");
      document.cookie = "hoge=hoge";
    },
    function () {
      // Storage access was denied.
      console.log("storage access was denied");
    }
  );
}

function setCookie() {
  document.cookie = "poko=poko";
}
