$(document).ready(function(){
    let votehtml = `
	<div class="vote border border-1 border-success rounded p-4 my-3">
	    <div class="form-group">
		<label for="VoteTitle">Vote Title</label>
		<textarea class="form-control" id="VoteTitle" rows="3"></textarea>
		<small id="emailHelp" class="form-text text-muted">Vote Title is just ask for sth.</small>
	    </div>
	    <div class="form-group">
		<label for="VoteType" style="display: block;">Vote Type</label>
		<div class="btn-group btn-group-sm mb-2" role="group" id="VoteType" aria-label="Basic radio toggle button group" style="user-select: auto;">
		    <input type="radio" class="btn-check" name="q" id="btnradio1" value="Radio" autocomplete="off" checked="" style="user-select: auto;">
		    <label class="btn btn-outline-dark" for="btnradio1" style="user-select: auto;">Radio</label>
		    <input type="radio" class="btn-check" name="q" id="btnradio2" value="CheckBox" autocomplete="off" style="user-select: auto;">
		    <label class="btn btn-outline-dark" for="btnradio2" style="user-select: auto;">CheckBox</label>
		</div>
	    </div>
	    <div class="input-group mb-3">
		<span class="input-group-text" id="basic-addon1">A</span>
		<input type="text" class="form-control" placeholder="Input content or leave blank" aria-label="A" aria-describedby="basic-addon1">
	    </div>
	    <div class="input-group mb-3">
		<span class="input-group-text" id="basic-addon2">B</span>
		<input type="text" class="form-control" placeholder="Input content or leave blank" aria-label="B" aria-describedby="basic-addon2">
	    </div>
	    <div class="input-group mb-3">
		<span class="input-group-text" id="basic-addon3">C</span>
		<input type="text" class="form-control" placeholder="Input content or leave blank" aria-label="C" aria-describedby="basic-addon3">
	    </div>
	    <div class="input-group mb-3">
		<span class="input-group-text" id="basic-addon4">D</span>
		<input type="text" class="form-control" placeholder="Input content or leave blank" aria-label="D" aria-describedby="basic-addon4">
	    </div>
	    <div class="input-group mb-3">
		<span class="input-group-text" id="basic-addon5">E</span>
		<input type="text" class="form-control" placeholder="Input content or leave blank" aria-label="E" aria-describedby="basic-addon5">
	    </div>
	    <div class="input-group mb-3">
		<span class="input-group-text" id="basic-addon6">F</span>
		<input type="text" class="form-control" placeholder="Input content or leave blank" aria-label="F" aria-describedby="basic-addon6">
	    </div>
	    <div class="input-group mb-3">
		<span class="input-group-text" id="basic-addon7">G</span>
		<input type="text" class="form-control" placeholder="Input content or leave blank" aria-label="G" aria-describedby="basic-addon7">
	    </div>
	    <div class="form-group">
		<label for="EnableTxtOpt" style="display: block;">Enable Text Option?</label>
		<div class="btn-group btn-group-sm mb-2" role="group" id="EnableTxtOpt" aria-label="Basic radio toggle button group" style="user-select: auto;">
		    <input type="radio" class="btn-check" name="e" id="btnEnable1" value="Yes" autocomplete="off" style="user-select: auto;">
		    <label class="btn btn-outline-dark" for="btnEnable1" style="user-select: auto;">Yes</label>
		    <input type="radio" class="btn-check" name="e" id="btnEnable2" value="No" autocomplete="off" checked="" style="user-select: auto;">
		    <label class="btn btn-outline-dark" for="btnEnable2" style="user-select: auto;">No</label>
		</div>
	    </div>
	</div>
    `
  $("#addvote").click(function(){
    $("#addvote").before(votehtml);
  });
});
