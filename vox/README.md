# vox wav files

Add the `vox/*.wav` files here, loop over and convert them to mp3, for example:

    for f in $(ls); do ffmpeg -y -i $f -af aresample=resampler=soxr -ar 44100 "$(basename $f .wav).mp3"; done

