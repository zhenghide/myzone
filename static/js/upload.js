$(function(){

    $('#uploadModal').on('show.bs.modal', function (event) {
        var button = $(event.relatedTarget);
        var photoId = button.data("id");
        var photoName = button.data('name');
        var modal = $(this);
        modal.find('.modal-title').text('上传照片到相册：' + photoName);
        modal.find('#photoId').html(photoId);
    });

    $('.pic-upload').fileinput({
        language: 'zh',
        uploadUrl: '/picture/upload',
        allowedFileExtensions : ['jpg', 'png','gif'],
        layoutTemplates :{
            //actionDelete:'', //去除上传预览的缩略图中的删除图标
            //actionUpload:'', //去除上传预览缩略图中的上传图片；
            actionZoom:'',   //去除上传预览缩略图中的查看详情预览的缩略图标。
        },
        uploadExtraData:function () {
            //向后台传递额外参数
            var photoId = $("#photoId").html();
            return {"photoId":photoId};
        }
    });

});

function closeModal() {
    window.location.href = "/photo";
}
