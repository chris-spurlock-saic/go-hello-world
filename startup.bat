CALL cls
CALL docker build -t chris_test_app . 
CALL docker run -it --rm -p 80:80 chris_test_app:latest